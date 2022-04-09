package record

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
