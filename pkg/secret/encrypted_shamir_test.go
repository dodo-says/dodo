package secret

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"filippo.io/age"
	"filippo.io/age/agessh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"testing"
)

func TestPoC_SplitEncryptDecryptCombine(t *testing.T) {

	privateKeyForAlice, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "generating private key")
	sshPublicKeyForAlice, err := ssh.NewPublicKey(&privateKeyForAlice.PublicKey)
	require.NoError(t, err, "generating ssh public key")

	privateKeyForBob, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "generating private key")
	sshPublicKeyForBob, err := ssh.NewPublicKey(&privateKeyForBob.PublicKey)
	require.NoError(t, err, "generating ssh public key")

	privateKeyForCharlie, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "generating private key")
	sshPublicKeyForCharlie, err := ssh.NewPublicKey(&privateKeyForCharlie.PublicKey)
	require.NoError(t, err, "generating ssh public key")

	secretMessage := "STRRL is a dodo"

	slices, err := SplitThenEncrypt([]byte(secretMessage), 3, 2, map[string][]byte{
		"alice":   ssh.MarshalAuthorizedKey(sshPublicKeyForAlice),
		"bob":     ssh.MarshalAuthorizedKey(sshPublicKeyForBob),
		"charlie": ssh.MarshalAuthorizedKey(sshPublicKeyForCharlie),
	})
	require.NoError(t, err, "splitting secret")

	aliceIdentity, err := agessh.NewRSAIdentity(privateKeyForAlice)
	require.NoError(t, err, "parse alice identity")

	bobIdentity, err := agessh.NewRSAIdentity(privateKeyForBob)
	require.NoError(t, err, "parse bob identity")

	charlieIdentity, err := agessh.NewRSAIdentity(privateKeyForCharlie)
	require.NoError(t, err, "parse charlie identity")
	sliceMap := make(map[string]EncryptedSlice)
	for _, slice := range slices {
		sliceMap[slice.Name] = slice
	}

	plainSliceReaderForAlice, err := age.Decrypt(bytes.NewReader(sliceMap["alice"].Content), aliceIdentity)
	require.NoError(t, err, "decrypting slice for alice")
	plainSliceForAlice, err := ioutil.ReadAll(plainSliceReaderForAlice)
	require.NoError(t, err, "reading slice for alice")

	plainSliceReaderForBob, err := age.Decrypt(bytes.NewReader(sliceMap["bob"].Content), bobIdentity)
	require.NoError(t, err, "decrypting slice for bob")
	plainSliceForBob, err := ioutil.ReadAll(plainSliceReaderForBob)
	require.NoError(t, err, "reading slice for bob")

	plainSliceReaderForCharlie, err := age.Decrypt(bytes.NewReader(sliceMap["charlie"].Content), charlieIdentity)
	require.NoError(t, err, "decrypting slice for charlie")
	plainSliceForCharlie, err := ioutil.ReadAll(plainSliceReaderForCharlie)
	require.NoError(t, err, "reading slice for charlie")

	message, err := Combine([][]byte{plainSliceForAlice, plainSliceForBob, plainSliceForCharlie})
	require.NoError(t, err, "combining slices")
	assert.Equalf(t, secretMessage, string(message), "combined slices")
}
