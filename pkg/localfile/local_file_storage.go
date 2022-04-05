package localfile

import (
	"context"
	"github.com/pkg/errors"
)

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

type MemberEntity struct {
	// The required unique name for this member. See Member.Name.
	Name string
	// The required name of which committee belongs to. See Member.CommitteeName.
	CommitteeName string
	// The optional description for this member. See Member.Description.
	Description string
	// The required public key for this member but base64-ed. See Member.PublicKey.
	PublicKeyBase64 string
}

type committeeMemberStorageModel struct {
	Length int            `json:"length"`
	Data   []MemberEntity `json:"data"`
}

func newCommitteeMemberStorageModel(length int, data []MemberEntity) *committeeMemberStorageModel {
	return &committeeMemberStorageModel{Length: length, Data: data}
}
func zeroValueCommitteeMemberStorageModel() *committeeMemberStorageModel {
	return newCommitteeMemberStorageModel(0, []MemberEntity{})
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

func (s *CommitteeMemberStorage) AddMember(ctx context.Context, member MemberEntity) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	storage.Data = append(storage.Data, member)
	return s.storage.write(ctx, *storage)
}

func (s *CommitteeMemberStorage) GetMember(ctx context.Context, committeeName string, memberName string) (*MemberEntity, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	for _, member := range storage.Data {
		if member.CommitteeName == committeeName && member.Name == memberName {
			return &member, nil
		}
	}
	return nil, errors.Errorf("member %s in committee %s not found", memberName, committeeName)
}

func (s *CommitteeMemberStorage) ListMemberInCommittee(ctx context.Context, committeeName string) ([]MemberEntity, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	var members []MemberEntity
	for _, member := range storage.Data {
		if member.CommitteeName == committeeName {
			members = append(members, member)
		}
	}
	return members, nil
}

func (s *CommitteeMemberStorage) ListMember(ctx context.Context) ([]MemberEntity, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	var members []MemberEntity
	for _, member := range storage.Data {
		members = append(members, member)
	}
	return members, nil
}
