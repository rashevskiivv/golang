package dump

import (
	"fmt"
	"unicode/utf8"
)

/*
A Go string is a read-only slice of bytes.
Strings are containers of text encoded in UTF-8.
Strings are equivalent to []byte.
In Go, the concept of a character is called a rune.
Rune is an integer that represents a Unicode code point.

To summarize, here are the salient points:

Go source code is always UTF-8.
A string holds arbitrary bytes.
A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
Those sequences represent Unicode code points, called runes.
No guarantee is made in Go that characters in strings are normalized.
*/
func strRunes() {
	const s = "สวัสดี"
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("\nRune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
