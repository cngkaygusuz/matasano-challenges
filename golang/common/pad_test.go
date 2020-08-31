package common

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Matasano 2.1
func Test_Pad_PKCS7(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")

	received := Pad_PKCS7(input, 4)
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	assert.Equal(t, expected, received)
}