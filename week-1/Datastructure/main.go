package main

import (
	"fmt"
)

// An array is a collection of elements of the same data type, stored under one variable name.
// Because arrays in Go are value types. Assigning one array to another creates a copy.
// Changing b doesn't change a.

// A slice is a dynamically sized, flexible view of an underlying array.

// This is one of the most important Go interview concepts.
// - len(slice) → number of elements currently present.
// - cap(slice) → maximum number of elements before a new underlying array is needed.

// A map is a data structure that stores data in key-value pairs.

// Think of a dictionary:
// - Key = Word
// - Value = Meaning

// A struct is a user-defined data type that groups multiple fields of different types into one structure.

type Userdetail struct {
	ID   int
	Name string
	age  int
}

func main() {
	var numbers [5]int

	// different way to create array
	// arr := [5]int{1, 2, 3, 4, 5}
	// value := [...]int{1, 2, 3, 4, 5, 3, 2}

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	// for i:=0;i<len(value);i++{
	// 	fmt.Println(i*value[i])
	// }

	// morden way use for loop
	// for index, val := range value {
	// 	fmt.Println(index * val)
	// }

	// fmt.Println(numbers)
	// fmt.Println(arr)
	// fmt.Println(len(value))

	// largest := 0
	// for i := 0; i < len(arr); i++ {
	// 	if arr[i] > largest {
	// 		largest = arr[i]
	// 	}
	// }

	// fmt.Println(largest)

	// slice  it is dynamic way to create array

	slicearr := make([]int, 5)

	for i := 0; i < len(slicearr); i++ {
		fmt.Println(i)
		slicearr[i] = (i + 1) * 2
	}

	fmt.Println(slicearr)

	// In slice make() take three parameters 1.[]_datatype 2. len 3. maxlen
	// append() - adding elements  it's not restricted you can extend the array
	// Index - adding element using index it is restricted show error

	slicearray := make([]int, 7, 8)

	for i := 0; i < len(slicearray); i++ {
		slicearray[i] = i * 2
	}
	// slicearray[3] = 30;
	slicearray = append(slicearray, 20)
	slicearray = append(slicearray, 20)
	slicearray = append(slicearray, 20)
	slicearray = append(slicearray, 20)

	fmt.Println(slicearray)

	// Map
	user := map[string]int{
		"name": 2,
		"age":  29,
	}

	// map using make
	// map[keytype]valuetype

	user2 := make(map[int]string)

	user2[1] = "name"
	user2[2] = "age"

	// store int and array of string
	user4 := make(map[int][]string)

	user4[0] = []string{"ff", "gg"}

	delete(user2, 1)

	fmt.Println(user, user2, user4)

	// struct

	// user5 := Userdetail{}

	// user5.Name = "Gg"
	// user5.age =12

	// second way to intialize value
	// user5 := Userdetail{
	// 	Name: "GG",
	// 	age:  12,
	// }

	// fmt.Println(user5)

	// Slice of Structs
	//Imagine an API returns multiple users. We can represent them using a slice of structs.

	users := []Userdetail{
		{ID: 1, Name: "Gourav", age: 23},
		{ID: 2, Name: "Rahul", age: 24},
		{ID: 3, Name: "Aman", age: 22},
	}

	for _, user := range users {
		fmt.Println(user.Name, user.age)
	}

	// Backend use case: When fetching users from a database, you commonly scan the results into a slice of structs.

	

}
