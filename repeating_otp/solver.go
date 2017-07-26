package repeating_otp

import (
	distance2 "github.com/cngkaygusuz/matasano-challenges/distance"
	"log"
	"github.com/cngkaygusuz/matasano-challenges/common"
	"github.com/cngkaygusuz/matasano-challenges/scotp/scoring"
)

const MIN_KEY_SIZE = 4
const MAX_KEY_SIZE = 60

func Solve(ciphertext []byte) (plaintext, key []byte ) {
	key_sizes := guess_key_size(ciphertext)

	var max_score int
	for _, ks := range key_sizes {
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
	bins := common.Bin(ciphertext)
	bins_transposed := common.Transpose(bins)

	transposed_solved := make(byte[][], len(bins_transposed))
	key = make(byte[], key_size)

	for i, bt := range bins_transposed {
		round_solution , round_key, round_score := scoring.Solve(bt)
		transposed_solved[i] = round_solution
		key[i] = round_key
		score += round_score
	}

	plaintext := common.Transpose(transposed_solved)
	return
}


func guess_key_size(ciphertext []byte) []int {
	retval := []int{-1, -1, -1}
	scores := []float64{1000000.0, 1000000.0, 1000000.0}

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
	first := ciphertext[0:key_size]
	second := ciphertext[key_size : key_size*2]
	distance := distance2.Hamming(first, second)
	return float64(distance) / float64(key_size)
}
