package scoring

import (
	"github.com/cngkaygusuz/matasano-challenges/common"
)

func Solve(ciphertext []byte) ([]byte, byte, int) {
	candidate := make([]byte, len(ciphertext))

	var best_guess []byte = make([]byte, len(ciphertext))
	var best_guess_score int = -2147483640
	var best_guess_key byte

	var i byte
	var score int
	for i = 0; i < 255; i++ {
		copy(candidate, ciphertext)
		score = 0

		for j := 0; j < len(candidate); j++ {
			candidate[j] = candidate[j] ^ i
			score += score_byte(candidate[j])
		}

		if score > best_guess_score {
			best_guess_score = score
			copy(best_guess, candidate)
			best_guess_key = i
		}

	}

	return best_guess, best_guess_key, best_guess_score
}

func score_byte(b byte) int {
	if !common.ReadableByte(b) { // case b is not readable
		return -10
	} else if b == 32 { // case b is space
		return 100
	} else if b == '\n' { // case b is newline character
		return 100
	} else if 65 <= b && b <= 90 { // case b is a capital alphabet character
		return 100
	} else if 97 <= b && b <= 122 { // case b is a lowercase alphabet character
		return 100
	} else { // case b is a readable character but not an alphabet character
		return 0
	}
}
