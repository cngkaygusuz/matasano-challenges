package common

func EqualBytes(first, second []byte) bool {
	length := len(first)

	if length != len(second) {
		return false
	}

	for i := 0; i < length; i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}