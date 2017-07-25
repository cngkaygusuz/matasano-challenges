package readable_bytes

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

	// Running this test prints the following:
	//
	// 2017/07/25 15:48:33 key: 0x47
	// 2017/07/25 15:48:33 plaintext: '\pptvqx?R\8l?svtz?~?opjq{?py?}~|pq'
	//
	// I speculate that this idea might have worked if the input set was bigger, but for smaller one, it seems it is
	// possible that we might get an all-readable ASCII output that is not the original plaintext.
}
