package main

import "fmt"

type Vertex struct {
	Num int
	Str string
}

type Tuple struct {
	FirstNum float64
	SecondNum float64
}

func main() {
	fmt.Println(Vertex{2, "Two"})

	vert := Vertex{4, "Four"}
	fmt.Println(vert.Str)
	accessStructByPointer()
}

func accessStructByPointer() {
	t := Tuple{5.3, 2.22}
	p := &t

	//(*p).SecondNum is an original syntax
	p.SecondNum *= 2
	fmt.Println(*p)
}