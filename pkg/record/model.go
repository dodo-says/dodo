package record

import "context"
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

	AddEncryptedRecord(ctx context.Context, encryptedRecord EncryptedRecordSlice) error
	GetEncryptedRecordsByRecordID(ctx context.Context, recordID uuid.UUID) ([]EncryptedRecordSlice, error)
	GetEncryptedRecordsByRecordIDAndMemberName(ctx context.Context, recordID uuid.UUID, memberName string) (EncryptedRecordSlice, error)

	CombineRecord(ctx context.Context, recordID uuid.UUID, slices []DecryptedRecordSlice) (DecryptedRecord, error)
}
