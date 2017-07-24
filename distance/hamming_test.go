package distance

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Hamming_1(t *testing.T) {
	first := []byte("this is a test")
	second := []byte("wokka wokka!!!")

	assert.Equal(t, 37, Hamming(first, second))
}