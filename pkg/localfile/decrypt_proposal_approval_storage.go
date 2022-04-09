package localfile

import (
	"context"

	"github.com/google/uuid"
)

type DecryptProposalApprovalEntity struct {
	ProposalID           uuid.UUID
	Member               string
	PlaintextSliceBase64 string
}

type decryptProposalApprovalStorageModel struct {
	Data []DecryptProposalApprovalEntity
}

func zeroValueDecryptProposalApprovalStorageModel() *decryptProposalApprovalStorageModel {
	return &decryptProposalApprovalStorageModel{
		Data: []DecryptProposalApprovalEntity{},
	}
}

type DecryptProposalApprovalStorage struct {
	storage *jsonFileStorage[decryptProposalApprovalStorageModel]
}

func NewDecryptProposalApprovalStorage(storagePath string) *DecryptProposalApprovalStorage {
	storage := newJsonFileStorage(storagePath, zeroValueDecryptProposalApprovalStorageModel)
	return &DecryptProposalApprovalStorage{
		storage: storage,
	}
}

func (s *DecryptProposalApprovalStorage) ListProposalApproval(ctx context.Context) ([]DecryptProposalApprovalEntity, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	return storage.Data, nil
}

func (s *DecryptProposalApprovalStorage) AddProposalApproval(ctx context.Context, proposal DecryptProposalApprovalEntity) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	storage.Data = append(storage.Data, proposal)
	return s.storage.write(ctx, *storage)
}
func (s *DecryptProposalApprovalStorage) CleanupProposalApprovalByProposalIDAndMember(ctx context.Context, proposalID uuid.UUID, member string) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	var filtered []DecryptProposalApprovalEntity
	for _, v := range storage.Data {
		if v.ProposalID != proposalID || v.Member != member {
			filtered = append(filtered, v)
		}
	}
	storage.Data = filtered
	return s.storage.write(ctx, *storage)
}
