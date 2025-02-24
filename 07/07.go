package main

import "fmt"

func IsValid(id int, text string) bool {
	return id > 0 && text != ""
}

func main() {
	// Таблица тестовых случаев
	tests := []struct {
		id       int
		text     string
		expected bool
	}{
		{1, "hello", true},  // id > 0, text не пустой
		{0, "hello", false}, // id <= 0, text не пустой
		{1, "", false},      // id > 0, text пустой
		{-5, "", false},     // id <= 0, text пустой
		{10, "world", true}, // id > 0, text не пустой
		{-1, "test", false}, // id <= 0, text не пустой
	}

	// Проходим по всем тестовым случаям
	for _, test := range tests {
		result := IsValid(test.id, test.text)

		// Выводим информацию о тесте
		fmt.Printf("Input: id = %d, text = %q\n", test.id, test.text)
		fmt.Printf("Expected: %v, Result: %v\n", test.expected, result)

		// Проверяем, совпадает ли результат с ожидаемым значением
		if result == test.expected {
			fmt.Println("✅ Test passed!")
		} else {
			fmt.Println("❌ Test failed!")
		}
	}
}
