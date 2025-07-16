/*
Как в Go реализуется наследование, если самого наследования нет?


Создай структуру `Animal` с методом `Speak()`, и структуру `Dog`, которая "наследует" поведение `Animal`, но переопределяет `Speak()`:

type Animal struct {}
func (a Animal) Speak() string { return "..." }

type Dog struct {
	Animal
}
func (d Dog) Speak() string { return "Woof" }
*/