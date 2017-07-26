package repeating_otp

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/require"
	"fmt"
	"encoding/hex"
)

func Test_Solve_1(t *testing.T) {
	input, err := ioutil.ReadFile("challenge-6-raw")
	require.Nil(t, err)

	plaintext, key, score := Solve(input)

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
