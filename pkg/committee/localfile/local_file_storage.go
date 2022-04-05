package localfile

import (
	"context"

	"github.com/dodo-says/dodo/pkg/committee"
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
	storage *jsonFileStorage[committeeStorageModel]
}

func NewCommitteeStorage(storagePath string) *CommitteeStorage {
	storage := newJsonFileStorage(storagePath, zeroValueCommitteeStorageModel)
	return &CommitteeStorage{
		storage: storage,
	}
}

func (s *CommitteeStorage) ListCommittee(ctx context.Context) ([]committee.Committee, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	return storage.Data, nil
}

func (s *CommitteeStorage) AddCommittee(ctx context.Context, committee committee.Committee) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	storage.Data = append(storage.Data, committee)
	return s.storage.write(ctx, *storage)
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
	storage *jsonFileStorage[committeeMemberStorageModel]
}

func NewCommitteeMemberStorage(storagePath string) *CommitteeMemberStorage {
	storage := newJsonFileStorage(storagePath, zeroValueCommitteeMemberStorageModel)
	return &CommitteeMemberStorage{
		storage: storage,
	}
}
