package ecb

import (
	"crypto/aes"
	"github.com/cngkaygusuz/matasano-challenges/common"
)

func Decrypt(ciphertext, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	common.PanicOnErr(err)

	retval := make([]byte, len(ciphertext) + 32)

	cipher.Decrypt(retval, ciphertext)

	return retval
}
