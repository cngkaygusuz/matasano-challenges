package common

import "log"

func Transpose(input [][]byte) [][]byte {
	if len(input) == 0 {
		return input
	}

	slice_size := len(input[0])
	for i := 1; i < len(input); i++ {
		if len(input[i]) != slice_size {
			log.Printf("each element in input should have the same length")
			log.Printf("first slice size: %d", slice_size)
			log.Printf("length of element %d: %d", i, len(input[i]))
		}
	}

	input_len := len(input)
	retval := make([][]byte, slice_size)
	for i := 0; i < slice_size ; i++ {
		retval[i] = make([]byte, input_len)
	}


	for i := 0; i < input_len; i++ {
		for j := 0; j < slice_size; j++ {
			retval[j][i] = input[i][j]
		}
	}

	return retval
}