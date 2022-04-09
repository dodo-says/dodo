package proposal

import "github.com/google/uuid"

type DecryptProposalApprove struct {
	RecordID       uuid.UUID
	Member         string
	PlaintextSlice []byte
}
