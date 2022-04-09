package localfile

import "context"

type committeeStorageModel struct {
	Length int               `json:"length"`
	Data   []CommitteeEntity `json:"data"`
}

type CommitteeEntity struct {
	Name        string
	Description string
}

func newCommitteeStorageModel(length int, data []CommitteeEntity) *committeeStorageModel {
	return &committeeStorageModel{Length: length, Data: data}
}
func zeroValueCommitteeStorageModel() *committeeStorageModel {
	return newCommitteeStorageModel(0, []CommitteeEntity{})
}

type CommitteeStorage struct {
	storage *jsonFileStorage[committeeStorageModel]
}

func NewCommitteeStorage(storagePath string) *CommitteeStorage {
	storage := newJsonFileStorage(storagePath, zeroValueCommitteeStorageModel)
	return &CommitteeStorage{
		storage: storage,
	}
}

func (s *CommitteeStorage) ListCommittee(ctx context.Context) ([]CommitteeEntity, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	return storage.Data, nil
}

func (s *CommitteeStorage) AddCommittee(ctx context.Context, committee CommitteeEntity) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	storage.Data = append(storage.Data, committee)
	return s.storage.write(ctx, *storage)
}
