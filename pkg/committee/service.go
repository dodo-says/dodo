package committee

import (
	"context"
	"github.com/dodo-says/dodo/pkg/localfile"
	"github.com/pkg/errors"
)

type Service interface {
	AddCommittee(ctx context.Context, committee Committee) error
	GetCommittee(ctx context.Context, committeeName string) (Committee, error)
	ListCommittee(ctx context.Context) ([]Committee, error)
	RemoveCommittee(ctx context.Context, committeeName string) error

	AddMemberToCommittee(ctx context.Context, committeeName string, member Member) error
	ListMemberOfCommittee(ctx context.Context, committeeName string) ([]Member, error)
	RemoveMemberFromCommittee(ctx context.Context, committeeName string, memberName string) error
	GetMemberOfCommittee(ctx context.Context, committeeName string, memberName string) (Member, error)
}

type ServiceImpl struct {
	committeeStorage *localfile.CommitteeStorage
	memberStorage    *localfile.CommitteeMemberStorage
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

func (s *ServiceImpl) GetCommittee(ctx context.Context, committeeName string) (Committee, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) ListCommittee(ctx context.Context) ([]Committee, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) RemoveCommittee(ctx context.Context, committeeName string) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) AddMemberToCommittee(ctx context.Context, committeeName string, member Member) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) ListMemberOfCommittee(ctx context.Context, committeeName string) ([]Member, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) RemoveMemberFromCommittee(ctx context.Context, committeeName string, memberName string) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceImpl) GetMemberOfCommittee(ctx context.Context, committeeName string, memberName string) (Member, error) {
	//TODO implement me
	panic("implement me")
}

func NewServiceImpl(committeeStorage *localfile.CommitteeStorage, memberStorage *localfile.CommitteeMemberStorage) *ServiceImpl {
	return &ServiceImpl{committeeStorage: committeeStorage, memberStorage: memberStorage}
}
