package common

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Memset(t *testing.T) {
	input := make([]byte, 5)
	expected := []byte{0, 0, 1, 1, 1}

	Memset(input, 2, uint8(1))

	assert.Equal(t, expected, input)
}

func Test_Memset_BadOffset(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	} ()

	input := make([]byte, 5)
	Memset(input, 10, 1)
}