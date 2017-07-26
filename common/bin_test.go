package common

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Binning a byte array of length 10 into length 3 bins results in 4 bin arrays ordered in a specific pattern.
func Test_Bin_BinIntoThree(t *testing.T) {
	input := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	received := Bin(input, 3)

	expected := [][]byte{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9},
	}

	assert.Equal(t, expected, received)
}

// Changing the input values does not change the received array
func Test_Bin_OriginalArrayChange(t *testing.T) {
	input := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	received := Bin(input, 3)

	expected := [][]byte{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9},
	}

	input[0] = 0xff

	assert.Equal(t, expected, received)
}
