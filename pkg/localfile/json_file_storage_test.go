package localfile

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestJsonFileStorage_write_then_read_string(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name: "empty",

			want:    "",
			wantErr: false,
		},
		{
			name:    "common",
			want:    "foo bar",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp(os.TempDir(), "json_file_storage_test")

			s := newJsonFileStorage[string](fmt.Sprintf("%s/%s", tempDir, "json_file.json"), func() *string {
				result := ""
				return &result
			})
			err = s.write(context.TODO(), tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := s.read(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonFileStorage_write_then_read_entity(t *testing.T) {
	type entity struct {
		Name        string
		Description string
	}
	t.Parallel()
	tests := []struct {
		name    string
		want    *entity
		wantErr bool
	}{
		{
			name: "empty",

			want:    &entity{},
			wantErr: false,
		},
		{
			name: "common",
			want: &entity{
				Name:        "foo",
				Description: "bar",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp(os.TempDir(), "json_file_storage_test")

			s := newJsonFileStorage[entity](fmt.Sprintf("%s/%s", tempDir, "json_file.json"), func() *entity {
				result := entity{}
				return &result
			})
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
