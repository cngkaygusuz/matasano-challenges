package solvers_common

import (
	"primitives"
	"strings"
)

func Decode(h primitives.Hex, key byte) string {
	newhex := h.Xor_bysingle(key)
	return string(newhex.Data)
}


func Makedict(dstr string) map[string]bool {
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

func Calculate_score(decodedstr *string, dictionary map[string]bool) int {
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
