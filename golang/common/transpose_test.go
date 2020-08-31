package common

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Transpose_SquareMatrix(t *testing.T) {
	input := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	received := Transpose(input)

	expected := [][]byte{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	assert.Equal(t, expected, received)
}

func Test_Transpose_RowMatrix(t *testing.T) {
	input := [][]byte{
		{1, 2, 3},
	}

	received := Transpose(input)

	expected := [][]byte{
		{1},
		{2},
		{3},
	}

	assert.Equal(t, expected, received)
}

func Test_Transpose_ColMatrix(t *testing.T) {
	input := [][]byte{
		{1},
		{2},
		{3},
	}

	received := Transpose(input)

	expected := [][]byte{
		{1, 2, 3},
	}

	assert.Equal(t, expected, received)
}

func Test_Transpose_NonSquareMatrix(t *testing.T) {
	input := [][]byte{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	received := Transpose(input)

	expected := [][]byte{
		{1, 3, 5},
		{2, 4, 6},
	}

	assert.Equal(t, expected, received)
}