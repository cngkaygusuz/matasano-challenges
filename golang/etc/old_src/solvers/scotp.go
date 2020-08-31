package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"primitives"
	"solvers_common"
)

// Single character one-time pad.
// This program takes a ciphertext encrypted with using only one byte and returns the plaintext.

// General method can be explained as; try all possible keys, score every output using a dictionary by just adding a
// fixed value if there are any words recognized and returning the output with the highest score.

func main() {
	// Matasano 1.3

	ciphertextPtr := flag.String("ciphertext", "", "Ciphertext to be broken.")
	dictpathPtr := flag.String("dictpath", "", "Path to the text file to be used as dictionary")

	flag.Parse()

	read, err := ioutil.ReadFile(*dictpathPtr)

	if err != nil {
		fmt.Println("Error reading the dictionary")
		panic(fmt.Sprint("Error reading the dictionary"))
	}

	dictionary := solvers_common.Makedict(string(read))

	var maxstr string
	var maxpoint int = -10  // Giving a negative value so the first decoded string will take this place, no matter score
	var maxkey int = -1

	for i:=0; i<256; i++ {
		key := byte(i)
		decodedstr := solvers_common.Decode(primitives.Hex_from_string(*ciphertextPtr), key)
		decodedpoint := solvers_common.Calculate_score(&decodedstr, dictionary)

		if decodedpoint > maxpoint {
			maxpoint = decodedpoint
			maxstr = decodedstr
			maxkey = i
		}

	}

	fmt.Printf("Best guess is '%s'\n", maxstr)
	fmt.Printf("Scored %d with key %d\n", maxpoint, maxkey)

}
