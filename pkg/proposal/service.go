package proposal

import (
	"context"
	"encoding/base64"
	"github.com/dodo-says/dodo/pkg/localfile"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service interface {
	CreateDecryptProposal(ctx context.Context, proposal DecryptProposal) error
	GetDecryptProposal(ctx context.Context, id uuid.UUID) (*DecryptProposal, error)
	ListDecryptProposal(ctx context.Context) ([]DecryptProposal, error)
	ListDecryptProposalByRecordID(ctx context.Context, recordID uuid.UUID) ([]DecryptProposal, error)

	CreateDecryptProposalApproval(ctx context.Context, proposal DecryptProposalApproval) error
	ListDecryptProposalApproval(ctx context.Context) ([]DecryptProposalApproval, error)
	ListDecryptProposalApprovalByProposalID(ctx context.Context, proposalID uuid.UUID) ([]DecryptProposalApproval, error)

	DecryptTheRecord(ctx context.Context, proposal DecryptProposal, approvals []DecryptProposalApproval) (*DecryptedRecord, error)
}

type ServiceImpl struct {
	proposalStorage         *localfile.DecryptProposalStorage
	proposalApprovalStorage *localfile.DecryptProposalApprovalStorage
}

func NewServiceImpl(proposalStorage *localfile.DecryptProposalStorage, proposalApprovalStorage *localfile.DecryptProposalApprovalStorage) *ServiceImpl {
	return &ServiceImpl{proposalStorage: proposalStorage, proposalApprovalStorage: proposalApprovalStorage}
}

func (s *ServiceImpl) CreateDecryptProposal(ctx context.Context, proposal DecryptProposal) error {
	err := s.proposalStorage.AddProposal(ctx, localfile.DecryptProposalEntity{
		ProposalID: proposal.ProposalID,
		RecordID:   proposal.RecordID,
		Reason:     proposal.Reason,
	})
	if err != nil {
		return errors.Wrap(err, "write proposal to storage")
	}
	return nil
}

func (s *ServiceImpl) GetDecryptProposal(ctx context.Context, id uuid.UUID) (*DecryptProposal, error) {
	proposals, err := s.ListDecryptProposal(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list proposal")
	}
	for _, proposal := range proposals {
		if proposal.ProposalID == id {
			return &proposal, nil
		}
	}
	return nil, errors.Errorf("proposal with id %s not found", id)
}

func (s *ServiceImpl) ListDecryptProposal(ctx context.Context) ([]DecryptProposal, error) {
	entities, err := s.proposalStorage.ListProposal(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "read proposal from storage")
	}
	proposals := make([]DecryptProposal, len(entities))
	for i, entity := range entities {
		proposals[i] = DecryptProposal{
			ProposalID: entity.ProposalID,
			RecordID:   entity.RecordID,
			Reason:     entity.Reason,
		}
	}
	return proposals, nil
}

func (s *ServiceImpl) ListDecryptProposalByRecordID(ctx context.Context, recordID uuid.UUID) ([]DecryptProposal, error) {
	proposals, err := s.ListDecryptProposal(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list proposal")
	}
	var filteredProposals []DecryptProposal
	for _, proposal := range proposals {
		if proposal.RecordID == recordID {
			filteredProposals = append(filteredProposals, proposal)
		}
	}
	return filteredProposals, nil
}

func (s *ServiceImpl) CreateDecryptProposalApproval(ctx context.Context, proposal DecryptProposalApproval) error {
	//TODO implement me
	panic("implement me")

}

func (s *ServiceImpl) ListDecryptProposalApproval(ctx context.Context) ([]DecryptProposalApproval, error) {
	approvalEntities, err := s.proposalApprovalStorage.ListProposalApproval(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "read proposal approval from storage")
	}
	var approvals []DecryptProposalApproval
	for _, approvalEntity := range approvalEntities {
		slice, err := base64.StdEncoding.DecodeString(approvalEntity.PlaintextSliceBase64)
		if err != nil {
			return nil, errors.Wrap(err, "decode plaintext slice")
		}
		approvals = append(approvals, DecryptProposalApproval{
			ProposalID:     approvalEntity.ProposalID,
			Member:         approvalEntity.Member,
			PlaintextSlice: slice,
		})
	}
	return approvals, nil
}

func (s *ServiceImpl) ListDecryptProposalApprovalByProposalID(ctx context.Context, proposalID uuid.UUID) ([]DecryptProposalApproval, error) {
	approvals, err := s.ListDecryptProposalApproval(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list proposal approval")
	}
	var filteredApprovals []DecryptProposalApproval
	for _, approval := range approvals {
		if approval.ProposalID == proposalID {
			filteredApprovals = append(filteredApprovals, approval)
		}
	}
	return filteredApprovals, nil
}

func (s *ServiceImpl) DecryptTheRecord(ctx context.Context, proposal DecryptProposal, approvals []DecryptProposalApproval) (*DecryptedRecord, error) {

	//TODO implement me
	panic("implement me")
}
