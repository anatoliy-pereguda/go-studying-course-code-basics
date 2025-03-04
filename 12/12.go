package main

import "fmt"

const (
	CodeOK = iota
	CodeCancelled
	CodeUnknown
)

func ErrorMessageToCode(msg string) int {
	switch msg {
	case "OK":
		return CodeOK
	case "CANCELLED":
		return CodeCancelled
	case "UNKNOWN":
		return CodeUnknown
	default:
		return CodeUnknown
	}
}

func main() {
	fmt.Println(ErrorMessageToCode("OK"))        // 0
	fmt.Println(ErrorMessageToCode("CANCELLED")) // 1
	fmt.Println(ErrorMessageToCode("UNKNOWN"))   // 2
	fmt.Println(ErrorMessageToCode("ANY"))       // 2
}
