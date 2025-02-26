package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = cases.Title(language.Russian).String(name)
	return fmt.Sprintf("Привет, %s!", name)
}

func main() {
	// Вызываем функцию Greetings с разными именами
	names := []string{
		"  анна  ",  // Имя с пробелами
		"МАРИЯ",     // Имя в верхнем регистре
		"екатерина", // Имя в нижнем регистре
	}

	for _, name := range names {
		result := Greetings(name)
		fmt.Println(result)
	}
}
