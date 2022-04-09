package proposal

import (
	"context"
	"github.com/dodo-says/dodo/pkg/localfile"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service interface {
	CreateDecryptProposal(ctx context.Context, proposal DecryptProposal) error
	ListDecryptProposal(ctx context.Context) ([]DecryptProposal, error)
	ListDecryptProposalByRecordID(ctx context.Context, recordID uuid.UUID) ([]DecryptProposal, error)

	CreateDecryptProposalApproval(ctx context.Context, proposal DecryptProposalApproval) error
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

func (s *ServiceImpl) ListDecryptProposalApprovalByProposalID(ctx context.Context, proposalID uuid.UUID) ([]DecryptProposalApproval, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) DecryptTheRecord(ctx context.Context, proposal DecryptProposal, approvals []DecryptProposalApproval) (*DecryptedRecord, error) {
	//TODO implement me
	panic("implement me")
}
