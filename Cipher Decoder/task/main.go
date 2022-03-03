package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var (
		m, n    string
		g       int
		v, x, z string
		p       int
	)
	_, err := fmt.Scan(&m, &n, &g, &v, &x, &z, &p)
	if err != nil {
		return
	}
	fmt.Println("OK")

	randomB := rand.Intn(p - 1)
	b := compute(g, randomB, p)

	var (
		u, i string
		a    int
	)
	_, err = fmt.Scan(&u, &i, &a)
	if err != nil {
		return
	}
	s := compute(a, randomB, p) % 26
	fmt.Printf("B is %v\n", b)

	mm := encrypt("Will you marry me?", s)

	fmt.Println(mm)

	var (
		t string
	)
	_, err = fmt.Scan(&t)
	if err != nil {
		return
	}

	ar := decrypt(t, s)

	if strings.Contains(ar, "Yeah,") {
		fmt.Println(encrypt("Great!", s))
	} else if strings.Contains(ar, "Let's") {
		fmt.Println(encrypt("What a pity!", s))
	}
}

func encrypt(text string, shiftNumber int) string { return cipher(text, 1, shiftNumber) }
func decrypt(text string, shiftNumber int) string { return cipher(text, -1, shiftNumber) }

func cipher(text string, direction int, shiftNumber int) string {
	// shift -> number of letters to move to right or left
	// offset -> size of the alphabet, in this case the plain ASCII
	shift, offset := rune(shiftNumber), rune(26)

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

func compute(a int, b int, p int) int {
	r := 1
	for b > 0 {
		r = r * a % p
		b--
	}
	return r
}
