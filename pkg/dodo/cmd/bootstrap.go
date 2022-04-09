package cmd

import (
	"fmt"
	"github.com/dodo-says/dodo/pkg/committee"
	"github.com/dodo-says/dodo/pkg/localfile"
	"github.com/dodo-says/dodo/pkg/record"
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

func BootstrapRecordService(storageDir string) record.Service {
	recordStorage := BootstrapRecordStorage(storageDir)
	encryptedRecordSliceStorage := BootstrapEncryptedRecordSliceStorage(storageDir)
	return record.NewServiceImpl(recordStorage, encryptedRecordSliceStorage)
}

func BootstrapRecordStorage(storageDir string) *localfile.RecordStorage {
	recordStorage := localfile.NewRecordStorage(fmt.Sprintf("%s/%s", storageDir, "record.json"))
	return recordStorage
}

func BootstrapEncryptedRecordSliceStorage(storageDir string) *localfile.EncryptedRecordSliceStorage {
	encryptedRecordSliceStorage := localfile.NewEncryptedRecordSliceStorage(fmt.Sprintf("%s/%s", storageDir, "encrypted_record_slice.json"))
	return encryptedRecordSliceStorage
}
