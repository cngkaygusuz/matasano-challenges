package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"strings"
	"primitives"
)

// Single character one-time pad.
// This program takes a ciphertext encrypted with using only one byte and returns the plaintext.

// General method can be explained as; try all possible keys, score every output using a dictionary by just summing
// the meaningful words and returning the output with the highest score.

func makedict(dstr string) map[string]bool {
	dictionary := make(map[string]bool)

	for _, el := range strings.Split(dstr, " ") {
		el = strings.Trim(el, "\n,!.?")
		el = strings.ToLower(el)

		if len(el) != 0 {
			dictionary[el] = true
		}
	}

	return dictionary
}

func decode(h primitives.Hex, key byte) string {
	newhex := h.Xor_bysingle(key)
	return string(newhex.Data)
}

func calculate_score(decodedstr *string, dictionary map[string]bool) int {
	decoded_split := strings.Split(*decodedstr, " ")
	points := 0

	for _, part := range decoded_split {
		part = strings.Trim(part, "\n,.!?")
		part = strings.ToLower(part)

		if dictionary[part] {
			points += 10
		}
	}

	return points
}

func main() {
	ciphertextPtr := flag.String("ciphertext", "", "Ciphertext to be broken.")
	dictpathPtr := flag.String("dictpath", "", "Path to the text file to be used as dictionary")

	flag.Parse()

	read, err := ioutil.ReadFile(*dictpathPtr)

	if err != nil {
		fmt.Println("Error reading the dictionary")
		panic(fmt.Sprint("Error reading the dictionary"))
	}

	dictionary := makedict(string(read))

	var maxstr string
	var maxpoint int = -10  // Giving a negative value so the first decoded string will take this place, no matter score
	var maxind int = -1

	for i:=0; i<256; i++ {
		key := byte(i)
		decodedstr := decode(primitives.Hex_from_string(*ciphertextPtr), key)
		decodedpoint := calculate_score(&decodedstr, dictionary)

		if decodedpoint > maxpoint {
			maxpoint = decodedpoint
			maxstr = decodedstr
			maxind = i
		}

	}

	fmt.Printf("Best guess is '%s'\n", maxstr)
	fmt.Printf("Scored %d at %d\n", maxpoint, maxind)

}
