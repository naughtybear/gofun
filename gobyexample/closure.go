package main

import (
	"fmt"
	"strings"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addZip := MakeAddSuffix(".zip")
	addTgz := MakeAddSuffix(".tar.gz")

	fmt.Println(addTgz("filename"), addZip("filename"), addZip("gobook.zip"))
}
