package record

import (
	"context"
	"github.com/dodo-says/dodo/pkg/committee"
	"github.com/dodo-says/dodo/pkg/localfile"
	"github.com/dodo-says/dodo/pkg/share"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service interface {
	BuildRecord(ctx context.Context, plainContent string, description string, committeeName string, threshold int) (*Record, []EncryptedRecordSlice, error)

	AddRecord(ctx context.Context, record Record) error
	GetRecord(ctx context.Context, id uuid.UUID) (Record, error)
	DeleteRecord(ctx context.Context, recordID uuid.UUID) error
	ListRecords(ctx context.Context) ([]Record, error)

	AddEncryptedRecordSlice(ctx context.Context, encryptedRecord EncryptedRecordSlice) error
	GetEncryptedRecordSlicesByRecordID(ctx context.Context, recordID uuid.UUID) ([]EncryptedRecordSlice, error)
	GetEncryptedRecordSliceByRecordIDAndMemberName(ctx context.Context, recordID uuid.UUID, memberName string) (EncryptedRecordSlice, error)

	CombineRecord(ctx context.Context, recordID uuid.UUID, slices []DecryptedRecordSlice) (DecryptedRecord, error)
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
			ID:         uuid.New(),
			RecordID:   recordId,
			MemberName: encryptedSlice.Name,
			Content:    encryptedSlice.Content,
		})
	}
	return &record, slices, nil
}

func (s *ServiceImpl) AddRecord(ctx context.Context, record Record) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) GetRecord(ctx context.Context, id uuid.UUID) (Record, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) DeleteRecord(ctx context.Context, recordID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) ListRecords(ctx context.Context) ([]Record, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) AddEncryptedRecordSlice(ctx context.Context, encryptedRecord EncryptedRecordSlice) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) GetEncryptedRecordSlicesByRecordID(ctx context.Context, recordID uuid.UUID) ([]EncryptedRecordSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) GetEncryptedRecordSliceByRecordIDAndMemberName(ctx context.Context, recordID uuid.UUID, memberName string) (EncryptedRecordSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) CombineRecord(ctx context.Context, recordID uuid.UUID, slices []DecryptedRecordSlice) (DecryptedRecord, error) {
	//TODO implement me
	panic("implement me")
}
