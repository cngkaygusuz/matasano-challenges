package primitives

import (
	"fmt"
)

type Hex struct {
	Data	[]byte
}

type ErrorCode int

func Hex_from_string(str string) Hex {
	// Construct a ByteStream object from given hexadecimal encoded string.
	// This method simply converts an ASCII representation of hexadecimal data to its true binary value
	// e.g
	// "dd" --> 0xdd
	// "a5" --> 0xa5

	if len(str) % 2 != 0 {
		str = "0" + str
	}

	hexlen := len(str) / 2

	var hexobj Hex;
	hexobj.Data = make([]byte, hexlen)

	for i := 0; i < hexlen; i++ {
		val1 := getHexValue(str[i*2])
		val2 :=	getHexValue(str[i*2+1])

		var b byte
		b = val1 << 4
		b = b | val2

		hexobj.Data[i] = b
	}
	return hexobj
}

func (b *Hex) Xor(other Hex) (Hex, ErrorCode) {
	if b.Len() != other.Len() {
		fmt.Printf("Lengths should be equal. Current %d, other %d", b.Len(), other.Len())
		var nullhex Hex
		return nullhex, 1
	}

	var newhex Hex
	newhex.Data = make([]byte, b.Len())

	for i, el := range b.Data {
		newhex.Data[i] = el ^ other.Data[i]
	}

	return newhex, 0

}

func (h *Hex) Len() int {
	return len(h.Data)
}

func (h *Hex) dump() string {
	return fmt.Sprintf("%x", h.Data)
}

func (h *Hex) Is_equal(other Hex) bool {
	if h.Len() != other.Len() {
		return false
	}

	for i, el := range h.Data {
		if el != other.Data[i] {
			return false
		}
	}

	return true
}

func (h *Hex) Xor_bysingle(key byte) Hex {
	var newhex Hex
	newhex.Data = make([]byte, len(h.Data))

	for i, el := range h.Data {
		newhex.Data[i] = el ^ key
	}

	return newhex
}
