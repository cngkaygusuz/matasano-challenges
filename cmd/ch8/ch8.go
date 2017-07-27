package main

import (
	"io/ioutil"
	"github.com/cngkaygusuz/matasano-challenges/common"
	"bytes"
	"log"
	"os"
)


func main() {
	input, err := ioutil.ReadFile("challenge-8.txt")
	common.PanicOnErr(err)

	lines := bytes.Split(input, []byte("\n"))

	for lineno, line := range lines {
		is_aes := detect_aes128(line)
		if is_aes {
			log.Printf("detected aes128 in line %d", lineno)
			os.Exit(0)
		}
	}

	log.Printf("couldn't detect anything")
}

func detect_aes128(input []byte) bool {
	bins := common.Bin(input, 16)

	for i := 0; i < len(bins)/2; i++ {
		for j := i+1; j < len(bins); j++ {
			if common.EqualBytes(bins[i], bins[j]) {
				return true
			}
		}
	}

	return false
}
