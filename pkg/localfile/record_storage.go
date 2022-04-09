package localfile

import (
	"context"
	"github.com/google/uuid"
)

type RecordEntity struct {
	ID            uuid.UUID `json:"id"`
	Description   string    `json:"description"`
	CommitteeName string    `json:"committee_name"`
	Threshold     int       `json:"threshold"`
}

type recordStorageModel struct {
	Data []RecordEntity `json:"data"`
}

func zeroValueRecordStorageModel() *recordStorageModel {
	return &recordStorageModel{
		Data: []RecordEntity{},
	}
}

type RecordStorage struct {
	storage *jsonFileStorage[recordStorageModel]
}

func NewRecordStorage(storagePath string) *RecordStorage {
	storage := newJsonFileStorage(storagePath, zeroValueRecordStorageModel)
	return &RecordStorage{
		storage: storage,
	}
}

func (s *RecordStorage) ListRecords(ctx context.Context) ([]RecordEntity, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	return storage.Data, nil
}

func (s *RecordStorage) AddRecord(ctx context.Context, record RecordEntity) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	storage.Data = append(storage.Data, record)
	return s.storage.write(ctx, *storage)
}
