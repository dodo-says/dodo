package localfile

import (
	"context"

	"github.com/google/uuid"
)

type DecryptProposalEntity struct {
	ProposalID uuid.UUID
	RecordID   uuid.UUID
	Reason     string
}

type decryptProposalStorageModel struct {
	Data []DecryptProposalEntity
}

func zeroValueDecryptProposalStorageModel() *decryptProposalStorageModel {
	return &decryptProposalStorageModel{
		Data: []DecryptProposalEntity{},
	}
}

type DecryptProposalStorage struct {
	storage *jsonFileStorage[decryptProposalStorageModel]
}

func NewDecryptProposalStorage(storagePath string) *DecryptProposalStorage {
	storage := newJsonFileStorage(storagePath, zeroValueDecryptProposalStorageModel)
	return &DecryptProposalStorage{
		storage: storage,
	}
}

func (s *DecryptProposalStorage) ListProposal(ctx context.Context) ([]DecryptProposalEntity, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	return storage.Data, nil
}

func (s *DecryptProposalStorage) AddProposal(ctx context.Context, proposal DecryptProposalEntity) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	storage.Data = append(storage.Data, proposal)
	return s.storage.write(ctx, *storage)
}
