package localfile

import (
	"context"
	"github.com/pkg/errors"
)

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
	Data []MemberEntity `json:"data"`
}

func newCommitteeMemberStorageModel(data []MemberEntity) *committeeMemberStorageModel {
	return &committeeMemberStorageModel{Data: data}
}
func zeroValueCommitteeMemberStorageModel() *committeeMemberStorageModel {
	return newCommitteeMemberStorageModel([]MemberEntity{})
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
