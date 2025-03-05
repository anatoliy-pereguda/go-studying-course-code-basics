package main

import "fmt"

// Определяем интерфейс Voicer
type Voicer interface {
	Voice() string
}

// Определяем структуру Cat
type Cat struct{}

func (c Cat) Voice() string {
	return "Мяу"
}

// Определяем структуру Cow
type Cow struct{}

func (c Cow) Voice() string {
	return "Мууу"
}

// Определяем структуру Dog
type Dog struct{}

func (d Dog) Voice() string {
	return "Гав"
}

func main() {
	// Создаем экземпляры структур
	cat := Cat{}
	dog := Dog{}
	cow := Cow{}

	// Вызываем метод Voice для каждого экземпляра
	fmt.Println(cat.Voice()) // Мяу
	fmt.Println(dog.Voice()) // Гав
	fmt.Println(cow.Voice()) // Мууу
}
