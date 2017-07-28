package repeating_otp

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/hex"
)

func Test_RepeatingOTP_1(t *testing.T) {
	cipher := Cipher{
		Plaintext: []byte{0, 0, 0},
		Key: []byte{1, 1, 1},
	}

	cipher.Encrypt()

	expected := []byte{1, 1, 1}

	assert.Equal(t, expected, cipher.Ciphertext)
}

func Test_RepeatingOTP_2(t *testing.T) {
	cipher := Cipher{
		Plaintext: []byte{0, 0, 0},
		Key: []byte{1},
	}

	cipher.Encrypt()

	expected := []byte{1, 1, 1}

	assert.Equal(t, expected, cipher.Ciphertext)
}

// Matasano 1.5
func Test_RepeatingOTP_Challenge(t *testing.T) {
	c1 := Cipher{
		Plaintext: []byte("Burning 'em, if you ain't quick and nimble"),
		Key: []byte("ICE"),
	}

	c1.Encrypt()

	assert.Equal(t, "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20", hex.EncodeToString(c1.Ciphertext))
}
