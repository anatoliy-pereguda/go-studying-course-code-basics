package main

import (
	"strings"
)

// Структура запроса
type UserCreateRequest struct {
	FirstName string // не может быть пустым; не может содержать пробелы
	Age       int    // не может быть 0 или отрицательным; не может быть больше 150
}

// Функция валидации
func Validate(req UserCreateRequest) string {
	// Проверка FirstName
	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return "invalid request"
	}

	// Проверка Age
	if req.Age <= 0 || req.Age > 150 {
		return "invalid request"
	}

	// Если все проверки пройдены, возвращаем пустую строку
	return "OK"
}

func main() {
	// Примеры использования
	println(Validate(UserCreateRequest{FirstName: "John", Age: 25}))     // ""
	println(Validate(UserCreateRequest{FirstName: "John Doe", Age: 25})) // "invalid request"
	println(Validate(UserCreateRequest{FirstName: "", Age: 25}))         // "invalid request"
	println(Validate(UserCreateRequest{FirstName: "John", Age: -5}))     // "invalid request"
	println(Validate(UserCreateRequest{FirstName: "John", Age: 200}))    // "invalid request"
}
