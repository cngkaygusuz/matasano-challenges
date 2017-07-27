package repeating_otp

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/require"
	"fmt"
	"encoding/hex"
)

// Matasano 6
func Test_Solve_1(t *testing.T) {
	input, err := ioutil.ReadFile("challenge-6-raw")
	require.Nil(t, err)

	plaintext, key, score := Solve(input, 5, 40)

	if plaintext == nil {
		fmt.Printf("plaintext is nil\n")
	} else if key == nil {
		fmt.Printf("key is nil\n")
	} else {
		fmt.Printf("score: %d\n", score)
		fmt.Printf("key length: %d\n", len(key))
		fmt.Printf("key: %s\n", hex.EncodeToString(key))
		fmt.Printf("key string: %s\n", string(key))
		fmt.Printf("%s\n", plaintext)
	}
}

func Test_Generate_Solve(t *testing.T) {
	plaintext := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	key := `Benim adim KIRMIZI`

	cipher := Cipher{
		Plaintext: []byte(plaintext),
		Key: []byte(key),
	}

	cipher.Encrypt()

	guessed_plaintext, guessed_key, score := Solve(cipher.Ciphertext, 5, 20)

	if guessed_plaintext == nil {
		fmt.Printf("plaintext is nil\n")
	} else if guessed_key == nil {
		fmt.Printf("key is nil\n")
	} else {
		fmt.Printf("score: %d\n", score)
		fmt.Printf("key length: %d\n", len(guessed_key))
		fmt.Printf("key: %s\n", hex.EncodeToString(guessed_key))
		fmt.Printf("key string: %s\n", string(guessed_key))
		fmt.Printf("%s\n", guessed_plaintext)
	}

}
