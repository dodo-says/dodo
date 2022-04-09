package localfile

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestLocalFileCommitteeMemberStorage_write_then_read(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *committeeMemberStorageModel
		wantErr bool
	}{
		{
			name: "empty",

			want: &committeeMemberStorageModel{
				Length: 0,
				Data:   []MemberEntity{},
			},
			wantErr: false,
		},
		{
			name: "empty with nil slice",
			want: &committeeMemberStorageModel{
				Length: 0,
				Data:   nil,
			},
			wantErr: false,
		},
		{
			name: "several",
			want: &committeeMemberStorageModel{
				Length: 2,
				Data: []MemberEntity{
					{
						Name:            "Alice",
						Description:     "Alice is the first committee",
						PublicKeyBase64: "Cg==",
					}, {
						Name:            "Bob",
						Description:     "Bob is the second committee",
						PublicKeyBase64: "c3NoLWVkMjU1MTkgQUFBQUMzTnphQzFsWkRJMU5URTVBQUFBSVBiWHhrM3FBSmM0VkVub3VLZ0FMSFUzMW1qRVlUYlZxVGdkYStENERrUDggc3RycmxAZ2l1Cg==",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp(os.TempDir(), "local_file_storage_test")

			s := NewCommitteeMemberStorage(fmt.Sprintf("%s/%s", tempDir, "committee_member.json"))
			err = s.storage.write(context.TODO(), *tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := s.storage.read(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("read() got = %v, want %v", got, tt.want)
			}
		})
	}
}
