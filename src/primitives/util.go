package primitives

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func isHex(b byte) bool {
	return ('0' <= b && b <= '9') || ('a' <= b && b <= 'f')
}

func getHexValue(b byte) byte {
	if '0' <= b && b <= '9' {
		return b - '0'
	} else {
		return b - 'a' + 10
	}
}
