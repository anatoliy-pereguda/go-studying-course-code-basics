package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// BEGIN (write your solution here)
func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = cases.Title(language.Russian).String(name)
	return fmt.Sprintf("Привет, %s!", name)
}

// END
