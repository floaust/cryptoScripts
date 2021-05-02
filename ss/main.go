package main

import (
	"github.com/pterm/pterm"
	"io/ioutil"
	"strings"
)

const text = `VJGTWF SMX VWE SHXWDTSMEW,
VWJ KWZJ KMWKKW TAJFWF LJMY,
ZAFY VWK XJMWZDAFYK DWLRLW HXDSMEW
MFV SF FMWKKWF FGUZ YWFMY.`

const validate = true

func main() {
	germanBytes, err := ioutil.ReadFile("language-german.lst.txt")
	englishBytes, err := ioutil.ReadFile("language-english.lst.txt")
	if err != nil {
		panic(err)
	}

	words := string(germanBytes) + "\n" + string(englishBytes)

	for i := 0; i < 26; i++ {
		t := cipher(text, 1, i) + "\n"
		t2 := cipher(text, -1, i)
		if validate {
			if strings.Contains(strings.ToLower(words), strings.ToLower(strings.Split(t, " ")[0])) {
				pterm.Printf("Offset: %d | Text: %s\n", i, t)
			}
			if strings.Contains(strings.ToLower(words), strings.ToLower(strings.Split(t2, " ")[0])) {
				pterm.Printf("Offset: %d | Text: %s\n", i, t2)
			}
		} else {
			pterm.Printf("Offset: %d | Text: %s\n", i, t)
			pterm.Printf("Offset: %d | Text: %s\n", i, t2)
		}

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
