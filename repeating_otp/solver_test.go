package repeating_otp

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/require"
	"encoding/base64"
)

func Test_Solve_1(t *testing.T) {
	input, err := ioutil.ReadFile("challenge-6.txt")
	require.Nil(t, err)

	input_decoded := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	base64.StdEncoding.Decode(input_decoded, input)

	guess_key_size(input_decoded)
}
