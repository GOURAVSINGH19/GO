package main

import (
	"fmt"
	"net/http"
)


func Homehandle(w http.ResponseWriter,r *http.Request){
	fmt.Println(w,"welcome to go backend");
}

func main() {
	a := 10
	b := 20
	fmt.Println("hello", a+b)

	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	result := <-ch
	fmt.Println(result)


	http.HandleFunc("/",Homehandle)
	http.ListenAndServe(":8000",nil);
}
