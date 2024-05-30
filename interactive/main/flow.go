package main

import (
	"fmt"
	"math"
	"errors"
)

func main() {
	fmt.Println(forLoop())
	fmt.Println(whileLike())

	res, err := sqrt(3.0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}

	shortConditionForm()

	upChar, err := charToUp("a")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Upcase char: ", upChar)
	}

	defferedExecutionFlow()
}

func forLoop() int {
	var sum = 0;

	for i := 0; i<10; i++ {
		sum += i
	}

	return sum
}

//For without init and transform works like C's while
func whileLike() int {
	var res int = 0;

	for res < 10 {
		res += 1
	} 

	return res
}

// func infiniteLoop() {
// 	for {
// 		fmt.Println("Infinite")
// 	}
// }

func sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, errors.New("Sqrt on negative numbers is prohibited")
	}

	return math.Sqrt(x), nil 
}

//Short if form
func shortConditionForm() {
	if res, err := sqrt(5); err == nil {
		fmt.Println("Result: ", res)
	}
}

func charToUp(c string) (string, error) {
	var res string
	
	switch c {
	case "a":
		res = "A"
	case "b":
		res = "B"
	case "c":
		res = "C"
	default:
		return "", errors.New("Mapping failed")
	}

	return res, nil
}

func defferedExecutionFlow() {
	defer fmt.Print("world")

	fmt.Print("hello ")
}