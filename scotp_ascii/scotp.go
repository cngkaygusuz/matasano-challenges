package scotp_ascii

import "github.com/cngkaygusuz/matasano-challenges/common"

func Solve(ciphertext []byte) ([]byte, byte) {
	candidate := make([]byte, len(ciphertext))

	var i byte
	for i = 0; i < 255; i++ {
		copy(candidate, ciphertext)
		for j := 0; j < len(candidate); j++ {
			candidate[j] = candidate[j] ^ i
		}

		if common.ReadableBytes(candidate) {
			return candidate, i
		}
	}

	return nil, 0
}
