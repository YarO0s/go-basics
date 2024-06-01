package main

import "fmt"

func main() {
	callable := func(x float64, y float64) float64 {
		return x * y
	}

	res := useCallable(callable)
	fmt.Println(res)

	//Closures
	positive, negative := simpleAdd(), simpleAdd()
	for i := 0; i < 10; i++ {
		fmt.Println("Positive: ", positive(i))
		fmt.Println("Negative: ", negative(-2 * i))
	}
}

func useCallable(fn func(float64, float64) float64) float64 {
	return fn(4.0, 2.0)	
}

func simpleAdd() func (x int) int {
	sum := 0
	return func (x int) int {
		sum += x
		return sum
	}
}

