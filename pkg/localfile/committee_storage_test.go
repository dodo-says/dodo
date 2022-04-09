package localfile

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestLocalFileCommitteeStorage_write_then_read(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *committeeStorageModel
		wantErr bool
	}{
		{
			name: "empty",

			want: &committeeStorageModel{
				Data: []CommitteeEntity{},
			},
			wantErr: false,
		},
		{
			name: "empty with nil slice",

			want: &committeeStorageModel{
				Data: nil,
			},
			wantErr: false,
		},
		{
			name: "several",
			want: &committeeStorageModel{
				Data: []CommitteeEntity{
					{
						Name:        "Alice",
						Description: "Alice is the first committee",
					}, {
						Name:        "Bob",
						Description: "Bob is the second committee",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp(os.TempDir(), "local_file_storage_test")

			s := NewCommitteeStorage(fmt.Sprintf("%s/%s", tempDir, "committee.json"))
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
