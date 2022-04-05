package cmd

import (
	"fmt"
	"github.com/dodo-says/dodo/pkg/committee"
	"github.com/dodo-says/dodo/pkg/localfile"
)

func BootstrapCommitteeService(storageDir string) committee.Service {
	committeeStorage := BootstrapCommitteeStorage(storageDir)
	memberStorage := BootstrapCommitteeMemberStorage(storageDir)
	return committee.NewServiceImpl(committeeStorage, memberStorage)
}

func BootstrapCommitteeMemberStorage(storageDir string) *localfile.CommitteeMemberStorage {
	memberStorage := localfile.NewCommitteeMemberStorage(fmt.Sprintf("%s/%s", storageDir, "committee_member.json"))
	return memberStorage
}

func BootstrapCommitteeStorage(storageDir string) *localfile.CommitteeStorage {
	committeeStorage := localfile.NewCommitteeStorage(fmt.Sprintf("%s/%s", storageDir, "committee.json"))
	return committeeStorage
}
