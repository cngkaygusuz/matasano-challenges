package primitives

import (
	"fmt"
)

// This is actually a constant.
var BASE64_ENCODINGS = map[byte]byte {
	0: 'A', 	16: 'Q',		32: 'g',		48: 'w',
	1: 'B',		17: 'R',		33: 'h',		49: 'x',
	2: 'C', 	18: 'S',		34: 'i',		50: 'y',
	3: 'D', 	19: 'T',		35: 'j',		51: 'z',
	4: 'E', 	20: 'U',		36: 'k',		52: '0',
	5: 'F', 	21: 'V',		37: 'l',		53: '1',
	6: 'G', 	22: 'W',		38: 'm',		54: '2',
	7: 'H', 	23: 'X',		39: 'n',		55: '3',
	8: 'I', 	24: 'Y',		40: 'o',		56: '4',
	9: 'J', 	25: 'Z',		41: 'p',		57: '5',
	10: 'K',	26: 'a',		42: 'q',		58: '6',
	11: 'L',	27: 'b',		43: 'r',		59: '7',
	12: 'M',	28: 'c',		44: 's',		60: '8',
	13: 'N',	29: 'd',		45: 't',		61: '9',
	14: 'O',	30: 'e',		46: 'u',		62: '+',
	15: 'P',	31: 'f',		47: 'v',		63: '/',
}

type Base64 struct {
	data	string
}

func Base64_from_hex(hexobj *Hex) Base64 {
	var b64obj Base64

	i := 0

	for i < hexobj.Len() {
		rb := hexobj.Data[i:i+3]
		readcnt := len(rb)

		//fmt.Println(bstr.readcnt)

		switch readcnt {
		case 0:
			goto breakloop
		case 1:
			fmt.Println("NOT IMPLEMENTED SOMETHING WILL GO BAD DO THIS NOW.")
		case 2:
			fmt.Println("NOT IMPLEMENTED SOMETHING WILL GO BAD DO THIS NOW.")
		case 3:
			var i1, i2, i3, i4 byte
			i1 = rb[0]
			i1 = i1 & 0xfc
			i1 = i1 >> 2
			b64obj.data += string(BASE64_ENCODINGS[i1])

			i2 = rb[0]
			i2 = i2 & 0x03
			i2 = i2 << 4
			i2 = i2 | (rb[1] & 0xf0) >> 4
			b64obj.data += string(BASE64_ENCODINGS[i2])

			i3 = rb[1]
			i3 = i3 & 0x0f
			i3 = i3 << 2
			i3 = i3 | (rb[2] & 0xc0) >> 6
			b64obj.data += string(BASE64_ENCODINGS[i3])

			i4 = rb[2]
			i4 = i4 & 0x3f
			b64obj.data += string(BASE64_ENCODINGS[i4])
		}
		i += 3
	}

	breakloop:
	return b64obj
}

