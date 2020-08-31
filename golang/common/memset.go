package common


func Memset(input []byte, offset int, value uint8) {
	length := len(input)

	if offset >= length {
		panic("offset is greater than or equal to input length")
	}

	if value == 0 {
		panic("value cannot be 0")
	}

	for i := offset; i < length; i++ {
		input[i] = value
	}
}