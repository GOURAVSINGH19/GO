package main

import (
	"fmt"
)

// 1. if-else if-else
// 2. switch (is used when you need to compare one value against multiple possible cases.)
// 3. switch case does not automatically fall through to the next case.
// 4. You usually don't need break after every case.
// 5. Unlike C++, Java, or JavaScript, Go does not have separate while or do-while loop keywords.
// 6. In GO, we can use for like a while loop
// 7. Go allows you to declare a variable inside an if statement.

func main() {
	age := 18
	// if-else if-else
	if age > 20 {
		fmt.Println("Age is greater than 18")
	} else if age == 18 {
		fmt.Println("Age is equal to 18")
	} else {
		fmt.Println("Age is lesser than 18")
	}

	// Go allows you to declare a variable inside an if statement.
	if value := 18; value < 20 {
		fmt.Println("Value is less than 20")
	} else {
		fmt.Println("Value is greater than 20")
	}

	// switch

	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	}

	// Go also supports condition-based switches.

	marks := 85
	switch {
	case marks >= 90:
		fmt.Println("A")
	case marks >= 75:
		fmt.Println("B")
	case marks >= 60:
		fmt.Println("C")
	default:
		fmt.Println("Fail")
	}

	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// use for loop like while loop
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// range loop
	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// If you only need the value:
	for _, value := range numbers {
		fmt.Println(value)
	}
}
