package primitives

import (
	"testing"
)

func Test_fromstring(t *testing.T) {
	expected_len := 3
	expected_data := [...]byte { 0xdd, 0xaf, 0xaf }

	var hexobj Hex
	hexobj = Hex_from_string("ddafaf")

	if hexobj.Len() != expected_len {
		t.Error("ERROR LENGTH")
	}

	for i := 0; i < expected_len; i++ {
		if expected_data[i] != hexobj.Data[i] {
			t.Errorf("Error; Expected %x, Got %x", expected_data[i], hexobj.Data[i])
		}
	}
}

func Test_xor(t *testing.T) {
	// Matasano 1.2
	hex1 := Hex_from_string("1c0111001f010100061a024b53535009181c")
	hex2 := Hex_from_string("686974207468652062756c6c277320657965")

	expected := Hex_from_string("746865206b696420646f6e277420706c6179")

	xorred, _ := hex1.Xor(hex2)

	if !xorred.Is_equal(expected) {
		t.Errorf("Unexpected output, produced and expected are printed respectively\n%s\n%s",
				 xorred.dump(), expected.dump())
	}
}
