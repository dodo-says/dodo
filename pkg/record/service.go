package record

import (
	"context"
	"encoding/base64"
	"github.com/dodo-says/dodo/pkg/committee"
	"github.com/dodo-says/dodo/pkg/localfile"
	"github.com/dodo-says/dodo/pkg/share"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service interface {
	BuildRecord(ctx context.Context, plainContent string, description string, committeeName string, threshold int) (*Record, []EncryptedRecordSlice, error)

	AddRecord(ctx context.Context, record Record) error
	GetRecord(ctx context.Context, id uuid.UUID) (*Record, error)
	DeleteRecord(ctx context.Context, recordID uuid.UUID) error
	ListRecords(ctx context.Context) ([]Record, error)
	ListRecordsByCommittee(ctx context.Context, committeeName string) ([]Record, error)

	AddEncryptedRecordSlice(ctx context.Context, encryptedRecord EncryptedRecordSlice) error
	ListEncryptedRecordSlices(ctx context.Context) ([]EncryptedRecordSlice, error)
	ListEncryptedRecordSlicesByRecordID(ctx context.Context, recordID uuid.UUID) ([]EncryptedRecordSlice, error)
	ListEncryptedRecordSliceByRecordIDAndMemberName(ctx context.Context, recordID uuid.UUID, memberName string) (*EncryptedRecordSlice, error)
}

type ServiceImpl struct {
	recordStorage               *localfile.RecordStorage
	encryptedRecordSliceStorage *localfile.EncryptedRecordSliceStorage
	committeeService            committee.Service
}

func NewServiceImpl(recordStorage *localfile.RecordStorage, encryptedRecordSliceStorage *localfile.EncryptedRecordSliceStorage, committeeService committee.Service) *ServiceImpl {
	return &ServiceImpl{recordStorage: recordStorage, encryptedRecordSliceStorage: encryptedRecordSliceStorage, committeeService: committeeService}
}

func (s *ServiceImpl) BuildRecord(ctx context.Context, plainContent string, description string, committeeName string, threshold int) (*Record, []EncryptedRecordSlice, error) {
	members, err := s.committeeService.ListMemberOfCommittee(ctx, committeeName)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "list members of committee %s", committeeName)
	}

	publicKeys := make(map[string][]byte)
	for _, member := range members {
		publicKeys[member.Name] = member.PublicKey
	}

	encryptedSlices, err := share.SplitThenEncrypt([]byte(plainContent), len(members), threshold, publicKeys)
	if err != nil {
		return nil, nil, errors.Wrap(err, "encrypt message")
	}

	recordId := uuid.New()
	record := Record{
		ID:            recordId,
		Description:   description,
		CommitteeName: committeeName,
		Threshold:     threshold,
	}

	var slices []EncryptedRecordSlice
	for _, encryptedSlice := range encryptedSlices {
		slices = append(slices, EncryptedRecordSlice{
			RecordID:   recordId,
			MemberName: encryptedSlice.Name,
			Content:    encryptedSlice.Content,
		})
	}
	return &record, slices, nil
}

func (s *ServiceImpl) AddRecord(ctx context.Context, record Record) error {
	err := s.recordStorage.AddRecord(ctx, localfile.RecordEntity{
		ID:            record.ID,
		Description:   record.Description,
		CommitteeName: record.CommitteeName,
		Threshold:     record.Threshold,
	})
	if err != nil {
		return errors.Wrap(err, "add record")
	}
	return nil
}

func (s *ServiceImpl) GetRecord(ctx context.Context, id uuid.UUID) (*Record, error) {
	records, err := s.ListRecords(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list records")
	}
	for _, record := range records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.Errorf("record with id %s not found", id)
}

func (s *ServiceImpl) DeleteRecord(ctx context.Context, recordID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) ListRecords(ctx context.Context) ([]Record, error) {
	recordEntities, err := s.recordStorage.ListRecords(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list records")
	}
	var records []Record
	for _, recordEntity := range recordEntities {
		records = append(records, Record{
			ID:            recordEntity.ID,
			Description:   recordEntity.Description,
			CommitteeName: recordEntity.CommitteeName,
			Threshold:     recordEntity.Threshold,
		})
	}
	return records, nil
}
func (s *ServiceImpl) ListRecordsByCommittee(ctx context.Context, committeeName string) ([]Record, error) {
	records, err := s.ListRecords(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list records")
	}
	var result []Record
	for _, record := range records {
		if record.CommitteeName == committeeName {
			result = append(result, record)
		}
	}
	return result, nil
}

func (s *ServiceImpl) AddEncryptedRecordSlice(ctx context.Context, encryptedRecord EncryptedRecordSlice) error {
	err := s.encryptedRecordSliceStorage.AddSlice(ctx, localfile.EncryptedRecordSliceEntity{
		RecordID:      encryptedRecord.RecordID,
		MemberName:    encryptedRecord.MemberName,
		ContentBase64: base64.StdEncoding.EncodeToString(encryptedRecord.Content),
	})
	if err != nil {
		return errors.Wrap(err, "add encrypted record slice")
	}
	return nil
}

func (s *ServiceImpl) ListEncryptedRecordSlices(ctx context.Context) ([]EncryptedRecordSlice, error) {
	slices, err := s.encryptedRecordSliceStorage.ListSlices(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list encrypted record slices")
	}
	var result []EncryptedRecordSlice
	for _, slice := range slices {
		base64Decode, err := base64.StdEncoding.DecodeString(slice.ContentBase64)
		if err != nil {
			return nil, errors.Wrap(err, "decode encrypted record slice")
		}
		result = append(result, EncryptedRecordSlice{
			RecordID:   slice.RecordID,
			MemberName: slice.MemberName,
			Content:    base64Decode,
		})
	}
	return result, nil
}

func (s *ServiceImpl) ListEncryptedRecordSlicesByRecordID(ctx context.Context, recordID uuid.UUID) ([]EncryptedRecordSlice, error) {
	slices, err := s.ListEncryptedRecordSlices(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "list encrypted record slices by record id %s", recordID)
	}
	var result []EncryptedRecordSlice
	for _, slice := range slices {
		if slice.RecordID == recordID {
			result = append(result, slice)
		}
	}
	return result, nil
}

func (s *ServiceImpl) ListEncryptedRecordSliceByRecordIDAndMemberName(ctx context.Context, recordID uuid.UUID, memberName string) (*EncryptedRecordSlice, error) {
	slices, err := s.ListEncryptedRecordSlicesByRecordID(ctx, recordID)
	if err != nil {
		return nil, errors.Wrapf(err, "list encrypted record slices by record id %s and member name %s", recordID, memberName)
	}
	for _, slice := range slices {
		slice := slice
		if slice.MemberName == memberName {
			return &slice, nil
		}
	}
	return nil, errors.Errorf("encrypted record slice not found, record id %s, member name %s", recordID, memberName)
}
