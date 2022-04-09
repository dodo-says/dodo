package proposal

import (
	"github.com/dodo-says/dodo/pkg/record"
	"github.com/google/uuid"
)

type DecryptProposal struct {
	ProposalID uuid.UUID
	RecordID   uuid.UUID
	Reason     string
}

type DecryptProposalApproval struct {
	ProposalID     uuid.UUID
	Member         string
	PlaintextSlice []byte
}

type DecryptedRecord struct {
	Proposal         DecryptProposal
	Record           record.Record
	CommitteeName    string
	ApprovedMembers  []string
	PlaintextContent []byte
}
