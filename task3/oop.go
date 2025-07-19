/*
Как в Go реализуется наследование, если самого наследования нет?
1. через embedding

Создай структуру `Animal` с методом `Speak()`, и структуру `Dog`, которая "наследует" поведение `Animal`, но переопределяет `Speak()`:

type Animal struct {}
func (a Animal) Speak() string { return "..." }

type Dog struct {
	Animal
}
func (d Dog) Speak() string { return "Woof" }
*/

package task3

import "fmt"

type Animal struct {
}

func (a Animal) Speak() string {
	return "..."
}

type Dog struct {
	Animal
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	animal := Animal{}
	dog := Dog{}

	fmt.Println(animal.Speak())
	fmt.Println(dog.Speak())
}
