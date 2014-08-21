package primitives

import (
	"testing"
//	"fmt"
)


func Test_6bytes(t *testing.T) {
	var hexobj Hex
	expected := "TWFuTWFu"

	hexobj.data = []byte("ManMan")

	b64 := base64_from_hex(&hexobj)

	if len(b64.data) != 8 {
		t.Errorf("Expected length 8, got %d", len(b64.data))
	}

	if b64.data != expected {
		t.Errorf("Error; produced and expected strings are printed respectively.\n%s\n%s", b64.data, expected)
	}
}

func Test_longbytes(t *testing.T) {
	// Matasano 1
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	bstr := hex_from_string(input)
	b64 := base64_from_hex(&bstr)

	if b64.data != expected {
		t.Errorf("Matasano challenge not met. Produced string and expected string given respectively.\n%s\n%s\n", b64.data, expected)
	}
}
