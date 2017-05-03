package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[123]"))
	fmt.Println(validID.MatchString("eve[8]"))
	fmt.Println(validID.MatchString("Job[234]"))
	fmt.Println(validID.MatchString("snakey"))

	fmt.Println(strings.Repeat("*", 20))

	// match string
	re := regexp.MustCompile("a.")
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllStringIndex("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllStringIndex("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllStringIndex("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
	fmt.Println(re.FindAllStringIndex("none", -1))

	fmt.Println(strings.Repeat("*", 20))

	// all string submatch
	re = regexp.MustCompile("a(x*)b")
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))

	fmt.Println(strings.Repeat("*", 20))

	// find string
	re = regexp.MustCompile("fo.?")
	fmt.Printf("%q\n", re.FindString("seafood"))
	fmt.Printf("%q\n", re.FindString("meat"))

	fmt.Println(strings.Repeat("*", 20))

	// find string index
	re = regexp.MustCompile("ab?")
	fmt.Println(re.FindStringIndex("tabblett"))
	fmt.Println(re.FindStringIndex("foo") == nil)

	fmt.Println(strings.Repeat("*", 20))

	// replace all literal string
	re = regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllLiteralString("--ab--axxb--", "T"))
	fmt.Println(re.ReplaceAllLiteralString("--ab--axxb--", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("--ab--axxb--", "${1}"))

	fmt.Println(strings.Repeat("*", 20))

	// replace all string
	fmt.Println(re.ReplaceAllString("--ab--axxb--", "T"))
	fmt.Println(re.FindAllString("--ab--axxb--", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("--ab--axxb--", -1))
	fmt.Println(re.ReplaceAllString("--ab--axxb--", "$1"))   // group {1} matches "" and "xx"
	fmt.Println(re.ReplaceAllString("--ab--axxb--", "${1}")) // 捕获组1分别匹配""和"xx", 所以 ab->"", axxb->"xx"
	fmt.Println(re.ReplaceAllString("--ab--axxb--", "${0}")) // 捕获组{0}是其本身, ab->"ab" axxb->"axxb"

	fmt.Println(strings.Repeat("*", 20))

	// split
	a := regexp.MustCompile("a")
	fmt.Println(a.Split("banana", -1))
	fmt.Println(a.Split("banana", 0))
	fmt.Println(a.Split("banana", 1))
	fmt.Println(a.Split("banana", 2))
	zp := regexp.MustCompile("z+")
	fmt.Println(zp.Split("pizza", -1))
	fmt.Println(zp.Split("pizza", 0))
	fmt.Println(zp.Split("pizza", 1))
	fmt.Println(zp.Split("pizza", 2))
	fmt.Println(zp.Split("pizza", 3))

	fmt.Println(strings.Repeat("*", 20))

	// subexpNames
	re = regexp.MustCompile("(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)")
	fmt.Println(re.MatchString("Alan Turing"))
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
	fmt.Println(re.ReplaceAllString("Alan Turing", "${last}"))
	fmt.Println(re.ReplaceAllString("Alan Turing S", "${last}"))
}
