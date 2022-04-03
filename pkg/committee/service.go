package committee

import "context"

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
