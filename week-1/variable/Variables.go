package main

import (
	"fmt"
	"strconv"
)

// IMP
// 1. Go doesn't automatically convert one into the other ( it is explict not implict language like c++)
// 2. we can use :=  outside the func
// 3. 

func main() {
	// var name string = "gourav"
	// var age int = 23

	// instead of using var we can use :=  (:= this is only used in func)

	name := "Gourav"
	age := 23

	// for const we have to use const keyword
	const money = 23
	// show compilation error
	//money= 34

	// conversion type

	// number to str Itao()
	newage := strconv.Itoa(age)

	// str to number use Atoi()
	newname, err := strconv.Atoi("234")

	if err != nil {
		fmt.Println("Invalid Number")
	} else {
		fmt.Println("str to number ", newname)
	}

	fmt.Println("Name", name)
	fmt.Println("number to str", newage)
	fmt.Println(money)
}
