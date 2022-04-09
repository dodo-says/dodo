package localfile

import (
	"context"
	"github.com/google/uuid"
)

type EncryptedRecordSliceEntity struct {
	ID            uuid.UUID `json:"id"`
	RecordID      uuid.UUID `json:"record_id"`
	MemberName    string    `json:"member_name"`
	ContentBase64 string    `json:"content_base64"`
}

type encryptedRecordSliceStorageModel struct {
	Data []EncryptedRecordSliceEntity `json:"data"`
}

func zeroValueEncryptedRecordSliceStorageModel() *encryptedRecordSliceStorageModel {
	return &encryptedRecordSliceStorageModel{
		Data: []EncryptedRecordSliceEntity{},
	}
}

type EncryptedRecordSliceStorage struct {
	storage *jsonFileStorage[encryptedRecordSliceStorageModel]
}

func NewEncryptedRecordSliceStorage(storagePath string) *EncryptedRecordSliceStorage {
	return &EncryptedRecordSliceStorage{
		storage: newJsonFileStorage(storagePath, zeroValueEncryptedRecordSliceStorageModel),
	}
}
func (s *EncryptedRecordSliceStorage) ListSlices(ctx context.Context) ([]EncryptedRecordSliceEntity, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	return storage.Data, nil
}
func (s *EncryptedRecordSliceStorage) AddRecord(ctx context.Context, slice EncryptedRecordSliceEntity) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	storage.Data = append(storage.Data, slice)
	return s.storage.write(ctx, *storage)
}
