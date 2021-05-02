package main

import "github.com/pterm/pterm"

const text = "VJGTWF SMX VWE SHXWDTSMEW"

func main() {
	for i := 0; i < 26; i++ {
		pterm.Println(cipher(text, -1, i))
	}

	for i := 0; i < 26; i++ {
		pterm.Println(cipher(text, 1, i))
	}
}

func cipher(text string, direction int, test int) string {
	// shift -> number of letters to move to right or left
	// offset -> size of the alphabet, in this case the plain ASCII
	shift, offset := rune(test), rune(26)

	// string->rune conversion
	runes := []rune(text)

	for index, char := range runes {
		// Iterate over all runes, and perform substitution
		// wherever possible. If the letter is not in the range
		// [1 .. 25], the offset defined above is added or
		// subtracted.
		switch direction {
		case -1: // encoding
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case +1: // decoding
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		// Above `if`s handle both upper and lower case ASCII
		// characters; anything else is returned as is (includes
		// numbers, punctuation and space).
		runes[index] = char
	}

	return string(runes)
}
