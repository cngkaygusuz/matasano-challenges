package scoring

import (
	"testing"
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"log"
)

// Matasano 3
func Test_Solve_1(t *testing.T) {
	input, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	require.Nil(t, err)

	plaintext, key := Solve(input)
	if plaintext == nil {
		log.Print("solver could not find a solution")
	} else {
		log.Printf("key: %#x", key)
		log.Printf("plaintext: '%s'", plaintext)
	}

	// Prints the following:
	//
	// Cooking MC's like a pound of bacon
}
