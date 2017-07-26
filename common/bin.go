package common


// Bin the input given the size
// This function copies the input array
func Bin(input []byte, bin_size int) [][]byte {
	if len(input) < bin_size {
		panic("length of input cannot be smaller than key size")
	}

	input_len := len(input)
	bin_length :=  ((input_len - 1) / bin_size) + 1

	bins := make([][]byte, bin_length)

	for i := 0; i < bin_length; i++ {
		bin := input[(i* bin_size):min(input_len, (i+1)*bin_size)]

		bin_copied := make([]byte, len(bin))
		copy(bin_copied, bin)

		bins[i] = bin_copied
	}

	return bins
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}