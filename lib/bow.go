package lib

import "fmt"

func Hi() string {
	fmt.Println("Hi")
	return "Hi"
}

func Bye() string {
	fmt.Println("Bye")
	return "Bye"
}

func Hello() string {
	fmt.Println("Hello")
	return "Hello"
}

func Goodbye() string {
	fmt.Println("Goodbye")
	return "Goodbye"
}

func Num(a, b, c int) (int, int, int) {
	return a + b, b * c, c / a
}
