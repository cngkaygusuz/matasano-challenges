package distance

// Hamming distance of two byte arrays, measured in bits.
func Hamming(first, second []byte) int {
	if len(first) != len(second) {
		panic("length of inputs does not match")
	}

	length := len(first)

	total_distance := 0

	for i := 0; i < length; i++ {
		diff := first[i] ^ second[i]

		for diff != 0 {
			total_distance += int(diff & 0x01)
			diff = diff >> 1
		}
	}

	return total_distance
}