package main

import "fmt"

func main() {
	a := 10
	p := &a

	//Pointer
	fmt.Println(p)
	//Value()
	fmt.Println(*p)

	//Accessing the value by a pointer
	*p = *p + 25
	fmt.Println(*p)
}