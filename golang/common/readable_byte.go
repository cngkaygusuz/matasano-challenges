package common

func ReadableByte(b byte) bool {
	// ASCII table, starting from decimal 32 to 126, consists of readable characters.
	if 32 <= b && b <= 126 {
		return true
	} else {
		return false
	}
}

func ReadableBytes(input []byte) bool {
	for _, b := range input {
		if !ReadableByte(b) {
			return false
		}
	}

	return true
}
