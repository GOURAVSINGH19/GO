package main 


// stack 
// 1. Stores function call frames and local execution data
// 2. Automatically grows and shrinks with calls
// 3. Typically short-lived data
// 4. Allocation is generally cheap


// heap

// 1. Stores dynamically managed objects
// 2. Managed by garbage collector
// 3. can outlive the function that created it
// 4. allocation and GC have additional costs

// to check value is store in heap or stack use command 
// go build -gcflags="-m" main.go


// result 
// ./main.go:33:6: can inline createUser
// ./main.go:39:26: inlining call to createUser
// ./main.go:40:16: inlining call to fmt.Println
// ./main.go:34:5: moved to heap: age
// ./main.go:40:16: ... argument does not escape
// ./main.go:40:17: *userAge escapes to heap



// Question

// age is declared inside createUser(). Why can we access it after the function returns?
// Because Go ensures that the value remains alive as long as it is needed. 
// In this case, the compiler's escape analysis may decide that age must be allocated on the heap.

//  what is escape analysis? 
// Escape analysis is a compiler process that determines whether a value can safely stay on the stack or needs to escape to the heap.


import "fmt"

func createUser() *int {
    age := 23
    return &age
}

func main() {
    userAge := createUser()
    fmt.Println(*userAge)
}