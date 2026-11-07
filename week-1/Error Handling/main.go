package main

import (
	"errors"
	"fmt"
)

// Error Handling

// Go generally handles errors by returning them as values rather than using exceptions.
// nil means there is no error 
// errors.New user for creates a fixed message
// fmt.Errorf() user for creates a formated  message

 
// painc and recover
// Use normal error returns for expected failures. Reserve panic/recover for exceptional situations, not routine business logic.


type error interface {
	Error() string
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	result,err := divide(20,10)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:-", result)
}
