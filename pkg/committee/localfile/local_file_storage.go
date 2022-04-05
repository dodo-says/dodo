package localfile

import (
	"context"
	"encoding/json"
	"github.com/dodo-says/dodo/pkg/committee"
	"github.com/pkg/errors"
	"os"
	"sync"
)

type committeeStorageModel struct {
	Length int                   `json:"length"`
	Data   []committee.Committee `json:"data"`
}

func newCommitteeStorageModel(length int, data []committee.Committee) *committeeStorageModel {
	return &committeeStorageModel{Length: length, Data: data}
}
func zeroValueCommitteeStorageModel() *committeeStorageModel {
	return newCommitteeStorageModel(0, []committee.Committee{})
}

type CommitteeStorage struct {
	// Filename of the file to storage data, example: "/tmp/tmp.4mc67VY2xs/committee.json".
	storagePath string
	// TODO: use flock
	rwlock sync.RWMutex
}

func NewCommitteeStorage(storagePath string) *CommitteeStorage {
	return &CommitteeStorage{storagePath: storagePath}
}

func (s *CommitteeStorage) read(ctx context.Context) (*committeeStorageModel, error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	stat, err := os.Stat(s.storagePath)

	// return zero value if file not exists
	if errors.Is(err, os.ErrNotExist) {
		return zeroValueCommitteeStorageModel(), nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "lookup committee storage file")
	}
	// return zero value if file is empty
	if stat.Size() == 0 {
		return zeroValueCommitteeStorageModel(), nil
	}

	// read content from file
	content, err := os.ReadFile(s.storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "read committee storage file")
	}
	result := zeroValueCommitteeStorageModel()
	err = json.Unmarshal(content, result)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal committee storage file")
	}
	return result, nil
}

func (s *CommitteeStorage) write(ctx context.Context, storageModel committeeStorageModel) error {
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

func (s *CommitteeStorage) ListCommittee(ctx context.Context) ([]committee.Committee, error) {
	storage, err := s.read(ctx)
	if err != nil {
		return nil, err
	}
	return storage.Data, nil
}

func (s *CommitteeStorage) AddCommittee(ctx context.Context, committee committee.Committee) error {
	storage, err := s.read(ctx)
	if err != nil {
		return err
	}
	storage.Data = append(storage.Data, committee)
	return s.write(ctx, *storage)
}

type memberEntity struct {
	// The required unique name for this member. See Member.Name.
	Name string
	// The optional description for this member. See Member.Description.
	Description string
	// The required public key for this member but base64-ed. See Member.PublicKey.
	PublicKeyBase64 string
}

type committeeMemberStorageModel struct {
	Length int            `json:"length"`
	Data   []memberEntity `json:"data"`
}

func newCommitteeMemberStorageModel(length int, data []memberEntity) *committeeMemberStorageModel {
	return &committeeMemberStorageModel{Length: length, Data: data}
}
func zeroValueCommitteeMemberStorageModel() *committeeMemberStorageModel {
	return newCommitteeMemberStorageModel(0, []memberEntity{})
}

type CommitteeMemberStorage struct {
	// Filename of the file to storage data, example: "/tmp/tmp.4mc67VY2xs/committee-member.json".
	storagePath string
	// TODO: use flock
	rwlock sync.RWMutex
}

func NewCommitteeMemberStorage(storagePath string) *CommitteeMemberStorage {
	return &CommitteeMemberStorage{storagePath: storagePath}
}

func (s *CommitteeMemberStorage) read(ctx context.Context) (*committeeMemberStorageModel, error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	stat, err := os.Stat(s.storagePath)

	// return zero value if file not exists
	if errors.Is(err, os.ErrNotExist) {
		return zeroValueCommitteeMemberStorageModel(), nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "lookup committee storage file")
	}
	// return zero value if file is empty
	if stat.Size() == 0 {
		return zeroValueCommitteeMemberStorageModel(), nil
	}

	// read content from file
	content, err := os.ReadFile(s.storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "read committee storage file")
	}
	result := zeroValueCommitteeMemberStorageModel()
	err = json.Unmarshal(content, result)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal committee storage file")
	}
	return result, nil
}

func (s *CommitteeMemberStorage) write(ctx context.Context, storageModel committeeMemberStorageModel) error {
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
