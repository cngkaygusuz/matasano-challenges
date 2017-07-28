package common


func Memset(input []byte, offset int, value uint8) {
	length := len(input)

	if offset >= length {
		panic("offset is greater than or equal to input length")
	}

	for i := offset; i < length; i++ {
		input[i] = value
	}
}