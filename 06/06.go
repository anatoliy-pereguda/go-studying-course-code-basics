package main

import (
	"fmt"
	"math"
)

func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	// Таблица тестовых случаев
	tests := []struct {
		x        int
		y        int
		expected int
	}{
		{3, 5, 3},      // x < y
		{10, 2, 2},     // x > y
		{7, 7, 7},      // x == y
		{-5, 10, -5},   // Отрицательное число меньше положительного
		{-10, -5, -10}, // Оба числа отрицательные
		{0, 0, 0},      // Оба числа равны нулю
	}

	// Проходим по всем тестовым случаям
	for _, test := range tests {
		result := MinInt(test.x, test.y)
		fmt.Printf("MinInt(%d, %d) = %d; expected = %d\n",
			test.x, test.y, result, test.expected)

		// Проверяем, совпадает ли результат с ожидаемым значением
		if result != test.expected {
			fmt.Println("❌ Test failed!")
		} else {
			fmt.Println("✅ Test passed!")
		}
	}
}
