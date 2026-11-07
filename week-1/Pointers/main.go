package main

import "fmt"

// Pointers are important for backend development because they help us work with memory and modify values through references.
// & Returns a variable's address.
// * Derefreances a pointer

// A method is a function attached to a particular type.
// In Go, we commonly attach methods to structs so that the struct can perform its own operations.

type User struct {
	ID   int
	Name string
	age  int
}

func (user *User) value() {
	user.age = 20
}

func main() {
	user := User{
		ID:   1,
		Name: "Gg",
		age:  2,
	}
	user.value()
	fmt.Println(user.age)

}
