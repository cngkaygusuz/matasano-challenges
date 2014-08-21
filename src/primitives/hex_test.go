package primitives

import (
	"testing"
)

func Test_fromstring(t *testing.T) {
	expected_len := 3
	expected_data := [...]byte { 0xdd, 0xaf, 0xaf }

	var hexobj Hex
	hexobj = hex_from_string("ddafaf")

	if hexobj.len() != expected_len {
		t.Error("ERROR LENGTH")
	}

	for i := 0; i < expected_len; i++ {
		if expected_data[i] != hexobj.data[i] {
			t.Errorf("Error; Expected %x, Got %x", expected_data[i], hexobj.data[i])
		}
	}
}

func Test_xor(t *testing.T) {
	// Matasano 2
	hex1 := hex_from_string("1c0111001f010100061a024b53535009181c")
	hex2 := hex_from_string("686974207468652062756c6c277320657965")

	expected := hex_from_string("746865206b696420646f6e277420706c6179")

	xorred, _ := hex1.xor(hex2)

	if !xorred.is_equal(expected) {
		t.Errorf("Unexpected output, produced and expected are printed respectively\n%s\n%s",
				 xorred.dump(), expected.dump())
	}
}
