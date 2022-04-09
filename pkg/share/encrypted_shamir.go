package share

import (
	"bytes"
	"filippo.io/age"
	"filippo.io/age/agessh"
	"github.com/hashicorp/vault/shamir"
	"github.com/pkg/errors"
)

// EncryptedSlice records the encrypted shamir share part for one.
type EncryptedSlice struct {
	// Name is the owner of public key.
	Name string
	// Content is the encrypted shamir share part.
	Content []byte
}

// SplitThenEncrypt would do the following:
// 1. Split the secret into n shares with shamir shares
// 2. Encrypt each share with the provided key
// 3. Return the encrypted shares
//
// secret is the data need to be encrypted; parts is the number of shares to be created; threshold is the minimum number of shares required to recover the secret;
// publicKeys is the map, key is the id of owner, value is the public key
func SplitThenEncrypt(secret []byte, parts int, threshold int, publicKeys map[string][]byte) ([]EncryptedSlice, error) {
	if len(publicKeys) != parts {
		return nil, errors.New("the number of public keys is not equal to the number of share parts")
	}

	slices, err := shamir.Split(secret, parts, threshold)
	if err != nil {
		return nil, errors.Wrapf(err, "split secret by shamir secret sharing, threshold: %d, parts: %d", threshold, parts)
	}

	var result []EncryptedSlice
	i := 0
	for name, publicKey := range publicKeys {
		slice := slices[i]
		recipient, err := agessh.ParseRecipient(string(publicKey))
		if err != nil {
			return nil, errors.Wrapf(err, "parse ssh public key for %s", name)
		}

		var content bytes.Buffer
		writeCloser, err := age.Encrypt(&content, recipient)
		if err != nil {
			return nil, errors.Wrapf(err, "prepare encrypt for %s", name)
		}
		_, err = writeCloser.Write(slice)
		if err != nil {
			return nil, errors.Wrapf(err, "encrypt slice for %s", name)
		}
		err = writeCloser.Close()
		if err != nil {
			return nil, errors.Wrapf(err, "close encrypt writer for %s", name)
		}
		result = append(result, EncryptedSlice{
			Name:    name,
			Content: content.Bytes(),
		})

		i += 1
	}

	return result, nil
}

// Combine would recover the secret from the encrypted shares and return the result.
func Combine(shares [][]byte) ([]byte, error) {
	combine, err := shamir.Combine(shares)
	if err != nil {
		return nil, errors.Wrap(err, "combine shares")
	}
	return combine, nil
}
