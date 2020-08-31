package common


func Pad_PKCS7(input []byte, pad_length uint8) []byte {
	retval := make([]byte, len(input) + int(pad_length))
	copy(retval, input)
	Memset(retval, len(input), pad_length)
	return retval
}
