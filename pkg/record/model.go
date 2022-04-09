package record

import (
	"context"
	"github.com/dodo-says/dodo/pkg/localfile"
)
import "github.com/google/uuid"

type Record struct {
	ID            uuid.UUID
	Description   string
	CommitteeName string
	Threshold     int
}

type EncryptedRecordSlice struct {
	ID         uuid.UUID
	RecordID   uuid.UUID
	MemberName string
	Content    []byte
}

type DecryptedRecordSlice struct {
	RecordID   uuid.UUID
	MemberName string
	Content    []byte
}

type DecryptedRecord struct {
	OriginalRecordID uuid.UUID
	CommitteeName    uuid.UUID
	Content          []byte
	ApprovedMembers  []string
}

type Service interface {
	BuildRecord(ctx context.Context, plainContent string, committeeName string, threshold int) (Record, []EncryptedRecordSlice, error)

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
}

func (s *ServiceImpl) BuildRecord(ctx context.Context, plainContent string, committeeName string, threshold int) (Record, []EncryptedRecordSlice, error) {
	//TODO implement me
	panic("implement me")
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

func NewServiceImpl(recordStorage *localfile.RecordStorage, encryptedRecordSliceStorage *localfile.EncryptedRecordSliceStorage) *ServiceImpl {
	return &ServiceImpl{recordStorage: recordStorage, encryptedRecordSliceStorage: encryptedRecordSliceStorage}
}
