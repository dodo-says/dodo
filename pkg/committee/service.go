package committee

import (
	"context"
	"encoding/base64"

	"github.com/dodo-says/dodo/pkg/localfile"
	"github.com/pkg/errors"
)

type Service interface {
	AddCommittee(ctx context.Context, committee Committee) error
	GetCommittee(ctx context.Context, committeeName string) (*Committee, error)
	ListCommittee(ctx context.Context) ([]Committee, error)
	RemoveCommittee(ctx context.Context, committeeName string) error

	AddMemberToCommittee(ctx context.Context, committeeName string, member Member) error
	ListMemberOfCommittee(ctx context.Context, committeeName string) ([]Member, error)
	RemoveMemberFromCommittee(ctx context.Context, committeeName string, memberName string) error
	GetMemberOfCommittee(ctx context.Context, committeeName string, memberName string) (*Member, error)
}

type ServiceImpl struct {
	committeeStorage *localfile.CommitteeStorage
	memberStorage    *localfile.CommitteeMemberStorage
}

func NewServiceImpl(committeeStorage *localfile.CommitteeStorage, memberStorage *localfile.CommitteeMemberStorage) *ServiceImpl {
	return &ServiceImpl{committeeStorage: committeeStorage, memberStorage: memberStorage}
}

func (s *ServiceImpl) AddCommittee(ctx context.Context, committee Committee) error {
	committees, err := s.committeeStorage.ListCommittee(ctx)
	if err != nil {
		return errors.Wrap(err, "list committee")
	}
	alreadyExist := false
	for _, c := range committees {
		if c.Name == committee.Name {
			alreadyExist = true
			break
		}
	}
	if alreadyExist {
		return errors.Errorf("committee %s already exist", committee.Name)
	}
	err = s.committeeStorage.AddCommittee(ctx, localfile.CommitteeEntity{
		Name:        committee.Name,
		Description: committee.Description,
	})
	if err != nil {
		return errors.Wrapf(err, "add committee %s", committee.Name)
	}
	return nil
}

func (s *ServiceImpl) GetCommittee(ctx context.Context, committeeName string) (*Committee, error) {
	entities, err := s.committeeStorage.ListCommittee(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list committee")
	}
	for _, e := range entities {
		if e.Name == committeeName {
			return &Committee{
				Name:        e.Name,
				Description: e.Description,
			}, nil
		}
	}
	return nil, errors.Errorf("committee %s not found", committeeName)
}

func (s *ServiceImpl) ListCommittee(ctx context.Context) ([]Committee, error) {
	entities, err := s.committeeStorage.ListCommittee(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list committee")
	}
	var committees []Committee
	for _, e := range entities {
		committees = append(committees, Committee{
			Name:        e.Name,
			Description: e.Description,
		})
	}
	return committees, nil
}

func (s *ServiceImpl) RemoveCommittee(ctx context.Context, committeeName string) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) AddMemberToCommittee(ctx context.Context, committeeName string, member Member) error {
	members, err := s.ListMemberOfCommittee(ctx, committeeName)
	if err != nil {
		return errors.Wrapf(err, "list member of committee %s", committeeName)
	}
	for _, m := range members {
		if m.Name == member.Name {
			return errors.Errorf("member %s already exist", member.Name)
		}
	}

	base64edPubKey := base64.StdEncoding.EncodeToString(member.PublicKey)

	entity := localfile.MemberEntity{
		Name:            member.Name,
		Description:     member.Description,
		CommitteeName:   committeeName,
		PublicKeyBase64: base64edPubKey,
	}
	err = s.memberStorage.AddMember(ctx, entity)
	if err != nil {
		return errors.Wrapf(err, "add member %s to committee %s", member.Name, committeeName)
	}
	return nil
}

func (s *ServiceImpl) ListMemberOfCommittee(ctx context.Context, committeeName string) ([]Member, error) {
	members, err := s.memberStorage.ListMemberInCommittee(ctx, committeeName)
	if err != nil {
		return nil, errors.Wrapf(err, "list member of committee %s", committeeName)
	}

	var result []Member
	for _, m := range members {

		pubKey, err := base64.StdEncoding.DecodeString(m.PublicKeyBase64)
		if err != nil {
			return nil, errors.Wrapf(err, "decode public key of member %s in committee %s", m.Name, committeeName)
		}
		result = append(result, Member{
			Name:        m.Name,
			Description: m.Description,
			PublicKey:   pubKey,
		})
	}
	return result, nil
}

func (s *ServiceImpl) RemoveMemberFromCommittee(ctx context.Context, committeeName string, memberName string) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) GetMemberOfCommittee(ctx context.Context, committeeName string, memberName string) (*Member, error) {
	//TODO implement me
	panic("implement me")
}
