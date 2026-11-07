package main

import "fmt"

// Interface :- It defines a set of method that a type must implement.
// in simple words interface decribes what something can do , not How it does it.

// Interfaces help us achieve abstraction and loose coupling.
//  Without interfaces, our code may depend directly on concrete implementations, making it harder to extend, test, and maintain.
// Interfaces allow different types to satisfy the same contract, so we can replace implementations without changing the business logic. They also make dependency injection and unit testing easier."
// One important thing to re
// member: Interfaces are not mandatory everywhere. For a small program with only one implementation, using an interface may add unnecessary complexity. Introduce interfaces when they provide a real benefit.

// Composition means building a complex type using smaller, reusable types.
// Go prefers composition over traditional class inheritance.

// interface
type Payment interface {
	Pay(amout int)
}

type UPI struct{}

func (u UPI) Pay(amount int) {
	fmt.Println("Paid using UPI", amount)
}

type Card struct{}

func (c Card) Pay(amount int) {
	fmt.Println("Paid using Card", amount)
}

func processPayment(p Payment) {
	p.Pay(1000.00)
}

// composition & embedding

type Person struct {
	Name string
	Age  int
}

type Department struct {
	Name string
}

// this is called embedding
type Employee struct {
	Person
	Department
	Salary int
}

// composition + interface

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}


// (c ConsoleLogger): Receiver. This method belongs to ConsoleLogger.
// Log(message string): Accepts a string.
// fmt.Println(message): Prints that string in the terminal.

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type UserService struct {
	logger Logger
}

func (s UserService) CreateUser() {
	s.logger.Log("User created successfully ")
}

func main() {

	upi := UPI{}

	card := Card{}

	processPayment(upi)
	processPayment(card)

	// composition

	emp := Employee{
		Person: Person{
			Name: "gg",
			Age:  23,
		},
		Department: Department{
			Name: "IT",
		},
		Salary: 50000,
	}

	fmt.Println(emp.Department.Name)
	fmt.Println(emp.Person.Name)

	// Composition + interface: step-by-step flow
	// 1. Create the concrete logger implementation (ConsoleLogger).
	// 2. Store it in the UserService logger field of type Logger.
	// 3. Go calls the ConsoleLogger.Log method through the Logger interface.
	userService := UserService{logger: ConsoleLogger{}}

	// 4. CreateUser triggers s.logger.Log("User created successfully ").
	userService.CreateUser()
}
