package localfile

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"sync"
)

type jsonFileStorage[T any] struct {
	// Filename of the file to storage data, example: "/tmp/tmp.4mc67VY2xs/committee.json".
	storagePath string
	// TODO: use flock
	rwlock               sync.RWMutex
	zeroValueConstructor func() *T
}

func newJsonFileStorage[T any](storagePath string, zeroValueConstructor func() *T) *jsonFileStorage[T] {
	return &jsonFileStorage[T]{storagePath: storagePath, zeroValueConstructor: zeroValueConstructor}
}

func (s *jsonFileStorage[T]) read(ctx context.Context) (*T, error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	stat, err := os.Stat(s.storagePath)

	// return zero value if file not exists
	if errors.Is(err, os.ErrNotExist) {
		return s.zeroValueConstructor(), nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "lookup committee storage file")
	}
	// return zero value if file is empty
	if stat.Size() == 0 {
		return s.zeroValueConstructor(), nil
	}

	// read content from file
	content, err := os.ReadFile(s.storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "read committee storage file")
	}
	result := s.zeroValueConstructor()
	err = json.Unmarshal(content, result)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal committee storage file")
	}
	return result, nil
}
func (s *jsonFileStorage[T]) write(ctx context.Context, data T) error {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	content, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "marshal committee storage file")
	}
	err = os.WriteFile(s.storagePath, content, 0o644)
	if err != nil {
		return errors.Wrap(err, "write committee storage file")
	}
	return nil
}
