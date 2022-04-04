package committee

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
		want    *localFileCommitteeStorageModel
		wantErr bool
	}{
		{
			name: "empty",

			want: &localFileCommitteeStorageModel{
				Length: 0,
				Data:   []Committee{},
			},
			wantErr: false,
		},
		{
			name: "empty with nil slice",

			want: &localFileCommitteeStorageModel{
				Length: 0,
				Data:   nil,
			},
			wantErr: false,
		},
		{
			name: "several",
			want: &localFileCommitteeStorageModel{
				Length: 2,
				Data: []Committee{
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

			s := NewLocalFileCommitteeStorage(fmt.Sprintf("%s/%s", tempDir, "local_file_storage.json"))
			err = s.write(context.TODO(), *tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := s.read(context.TODO())
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
