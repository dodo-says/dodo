package localfile

import (
	"context"
	"encoding/base64"

	"github.com/dodo-says/dodo/pkg/committee"
	"github.com/pkg/errors"
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
	// The required name of which committee belongs to. See Member.CommitteeName.
	CommitteeName string
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

func (s *CommitteeMemberStorage) AddMember(ctx context.Context, member committee.Member) error {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return err
	}
	base64edPublicKey := base64.StdEncoding.EncodeToString(member.PublicKey)
	storage.Data = append(storage.Data, memberEntity{
		Name:            member.Name,
		CommitteeName:   member.CommitteeName,
		Description:     member.Description,
		PublicKeyBase64: base64edPublicKey,
	})
	return s.storage.write(ctx, *storage)
}

func (s *CommitteeMemberStorage) GetMember(ctx context.Context, committeeName string, memberName string) (*committee.Member, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	for _, member := range storage.Data {
		if member.CommitteeName == committeeName && member.Name == memberName {
			publicKey, err := base64.StdEncoding.DecodeString(member.PublicKeyBase64)
			if err != nil {
				return nil, errors.Wrap(err, "decode public key")
			}
			return &committee.Member{
				Name:          member.Name,
				CommitteeName: member.CommitteeName,
				Description:   member.Description,
				PublicKey:     publicKey,
			}, nil
		}
	}
	return nil, errors.Errorf("member %s in committee %s not found", memberName, committeeName)
}

func (s *CommitteeMemberStorage) ListMemberInCommittee(ctx context.Context, committeeName string) ([]committee.Member, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}

	var members []committee.Member
	for _, member := range storage.Data {
		if member.CommitteeName == committeeName {
			publicKey, err := base64.StdEncoding.DecodeString(member.PublicKeyBase64)
			if err != nil {
				return nil, errors.Wrap(err, "decode public key")
			}
			members = append(members, committee.Member{
				Name:          member.Name,
				CommitteeName: member.CommitteeName,
				Description:   member.Description,
				PublicKey:     publicKey,
			})
		}
	}
	return members, nil
}

func (s *CommitteeMemberStorage) ListMember(ctx context.Context) ([]committee.Member, error) {
	storage, err := s.storage.read(ctx)
	if err != nil {
		return nil, err
	}
	var members []committee.Member
	for _, member := range storage.Data {
		publicKey, err := base64.StdEncoding.DecodeString(member.PublicKeyBase64)
		if err != nil {
			return nil, errors.Wrap(err, "decode public key")
		}
		members = append(members, committee.Member{
			Name:          member.Name,
			CommitteeName: member.CommitteeName,
			Description:   member.Description,
			PublicKey:     publicKey,
		})
	}
	return members, nil
}
