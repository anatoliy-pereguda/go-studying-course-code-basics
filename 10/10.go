package main

import (
	"strings"
)

func ModifySpaces(s, mode string) string {
	switch mode {
	case "dash":
		return strings.ReplaceAll(s, " ", "-")
	case "underscore":
		return strings.ReplaceAll(s, " ", "_")
	case "unknown":
		return strings.ReplaceAll(s, " ", "*")
	case "":
		return strings.ReplaceAll(s, " ", "*")
	default:
		return strings.ReplaceAll(s, " ", " ")
	}
}

func main() {
	println(ModifySpaces("hello world", "dash"))       // hello-world
	println(ModifySpaces("hello world", "underscore")) // hello_world
	println(ModifySpaces("hello world", "unknown"))    // hello*world
	println(ModifySpaces("hello world", ""))           // hello*world
	println(ModifySpaces("hello world", "any"))        // hello world
}
