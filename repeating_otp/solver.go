package repeating_otp

import (
	distance2 "github.com/cngkaygusuz/matasano-challenges/distance"
	"log"
	"github.com/cngkaygusuz/matasano-challenges/common"
	"github.com/cngkaygusuz/matasano-challenges/scotp/scoring"
)

const MIN_KEY_SIZE = 4
const MAX_KEY_SIZE = 40

func Solve(ciphertext []byte) (plaintext, key []byte, max_score int) {
	key_sizes := guess_key_size(ciphertext)

	max_score = -2147483640
	for _, ks := range key_sizes {
		if ks == -1 {
			continue
		}

		candidate_plaintext, candidate_key, score := solve_for_keysize(ciphertext, ks)

		if score > max_score {
			max_score = score
			plaintext = candidate_plaintext
			key = candidate_key
		}
	}

	return
}

func solve_for_keysize(ciphertext []byte, key_size int) (plaintext, key []byte, score int) {
	bins := common.Bin(ciphertext, key_size)

	if len(bins[len(bins)-1]) != key_size {
		bins = bins[0:len(bins)-1]
	}

	bins_transposed := common.Transpose(bins)

	transposed_solved := make([][]byte, len(bins_transposed))
	key = make([]byte, key_size)

	for i, bt := range bins_transposed {
		round_solution , round_key, round_score := scoring.Solve(bt)
		transposed_solved[i] = round_solution
		key[i] = round_key
		score += round_score
	}

	plaintext_binned := common.Transpose(transposed_solved)
	plaintext = common.Unbin(plaintext_binned)
	return
}


func guess_key_size(ciphertext []byte) []int {
	retval := []int{-1, -1, -1, -1 ,-1, -1}
	scores := []float64{1000000.0, 1000000.0, 1000000.0, 1000000.0, 1000000.0, 1000000.0}

	for key_size := MIN_KEY_SIZE; key_size <= MAX_KEY_SIZE; key_size++ {
		current_score := get_score_for_size(ciphertext, key_size)
		ind := get_biggest_smaller_than(current_score, scores)
		if ind != -1 {
			scores[ind] = current_score
			retval[ind] = key_size
		}
	}

	log.Printf("key sizes: %#v", retval)
	log.Printf("scores: %#v", scores)

	return retval
}

func get_biggest_smaller_than(current_score float64, input []float64, ) int {
	ind := -1
	diff := 100000000.0

	for i, score := range input {
		if current_score > score {
			continue
		}

		current_diff := score - current_score
		if current_diff < diff {
			diff = current_diff
			ind = i
		}
	}

	return ind
}

func get_score_for_size(ciphertext []byte, key_size int) float64 {
	total_hamming_distance := 0
	bins := common.Bin(ciphertext, key_size)
	for i := 0; i < len(bins) / 2; i++ {
		total_hamming_distance += distance2.Hamming(bins[i], bins[i*2])
	}

	return float64(total_hamming_distance) / (float64(key_size) * float64(len(bins)))
}
