package ecb

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/require"
	"log"
)


// Matasano 7
func Test_Decrypt_1(t *testing.T) {
	input, err := ioutil.ReadFile("challenge-7-raw")
	require.Nil(t, err)

	key := []byte("YELLOW SUBMARINE")

	plaintext := Decrypt(input, key)

	log.Print(string(plaintext))
}