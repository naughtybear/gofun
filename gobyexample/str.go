package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	text1 := "\"what's that?\", he said"
	text2 := `"what's that?", he said`
	fmt.Println(text1)
	fmt.Println(text2)

	fmt.Printf("%t %t\n", true, false)
	fmt.Printf("%T %T\n", true, false)

	fmt.Printf("%d %x %#04x %U '%c'\n", 0x3A6, 934, 934, '\u03A6', '\U000003A6')

	for _, x := range []float64{-.258, 7194.84, -60897162.0218, 1.5000089e-8} {
		fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ','))
	}

	fmt.Println(strings.IndexFunc("s ", unicode.IsSpace))
}

func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:] // remove "+/-"
	log.Println(whole)
	fraction := ""
	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:] // remove "+/-0"
	}
	sep := string(separator)
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	if amount < 0.0 {
		whole = "-" + whole
	}
	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}
