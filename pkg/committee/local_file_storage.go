package committee

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"sync"
)

type localFileCommitteeStorageModel struct {
	Length int         `json:"length"`
	Data   []Committee `json:"data"`
}

func newLocalFileCommitteeStorageModel(length int, data []Committee) *localFileCommitteeStorageModel {
	return &localFileCommitteeStorageModel{Length: length, Data: data}
}
func zeroLocalFileCommitteeStorageModel() *localFileCommitteeStorageModel {
	return newLocalFileCommitteeStorageModel(0, []Committee{})
}

type LocalFileCommitteeStorage struct {
	// Filename of the file to storage data, example: "/tmp/tmp.4mc67VY2xs/committee.json".
	storagePath string
	// TODO: use flock
	rwlock sync.RWMutex
}

func NewLocalFileCommitteeStorage(storagePath string) *LocalFileCommitteeStorage {
	return &LocalFileCommitteeStorage{storagePath: storagePath}
}

func (s *LocalFileCommitteeStorage) read(ctx context.Context) (*localFileCommitteeStorageModel, error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	stat, err := os.Stat(s.storagePath)

	// return zero value if file not exists
	if errors.Is(err, os.ErrNotExist) {
		return zeroLocalFileCommitteeStorageModel(), nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "lookup committee storage file")
	}
	// return zero value if file is empty
	if stat.Size() == 0 {
		return zeroLocalFileCommitteeStorageModel(), nil
	}

	// read content from file
	content, err := os.ReadFile(s.storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "read committee storage file")
	}
	result := zeroLocalFileCommitteeStorageModel()
	err = json.Unmarshal(content, result)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal committee storage file")
	}
	return result, nil
}

func (s *LocalFileCommitteeStorage) write(ctx context.Context, storageModel localFileCommitteeStorageModel) error {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	content, err := json.Marshal(storageModel)
	if err != nil {
		return errors.Wrap(err, "marshal committee storage file")
	}
	err = os.WriteFile(s.storagePath, content, 0o644)
	if err != nil {
		return errors.Wrap(err, "write committee storage file")
	}
	return nil
}

func (s *LocalFileCommitteeStorage) ListCommittee(ctx context.Context) ([]Committee, error) {
	panic("not implemented")
}

type LocalFileCommitteeMemberStorage struct {
}
