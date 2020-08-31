package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"strings"
	"solvers_common"
	"primitives"
)


// Single Character One-Time Pad using Multiple Entries
// This is same as "scotp", but instead trying to break only one entry, we will try to find a meaningful text encrypted
// with a single byte key amongst multiple values.

func main() {
	// Matasano 1.3

	ciphertextspathPtr := flag.String("ctextpath", "", "Ciphertext to be broken.")
	dictpathPtr := flag.String("dictpath", "", "Path to the text file to be used as dictionary")

	flag.Parse()

	read_dict, err_dict := ioutil.ReadFile(*dictpathPtr)
	if err_dict != nil {
		panic(fmt.Sprint("Error reading the dictionary"))
	}
	dictionary := solvers_common.Makedict(string(read_dict))

	read_cph, err_cph := ioutil.ReadFile(*ciphertextspathPtr)
	if err_cph != nil {
		panic(fmt.Sprint("Error reading the dictionary"))
	}

	ciphertexts := strings.Split(string(read_cph), "\n")

	var maxstr string
	var maxpoint int = -10
	var maxkey int
	var maxind int

	for ctextind, ciphertext := range ciphertexts {
		for i:=0; i<256; i++ {
			key := byte(i)
			decodedstr := solvers_common.Decode(primitives.Hex_from_string(ciphertext), key)
			decodedpoint := solvers_common.Calculate_score(&decodedstr, dictionary)

			if decodedpoint > maxpoint {
				maxpoint = decodedpoint
				maxstr = decodedstr
				maxkey = i
				maxind = ctextind
			}
		}
	}

	fmt.Printf("Guessed string is: %s", maxstr)
	fmt.Printf("Found at %d with key of %d and score of %d", maxind, maxkey, maxpoint)

}
