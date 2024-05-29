package main

//factorized import, substitution for import fmt; import math 
import (
	"fmt"
	"math"
)

//Package level vars
//Var cannot be omitted
var packVarOne, packVarTwo bool = true, true;

//Constants
const A = 34

func main() {
	//Exported functions
	exported()

	//Typed args
	fmt.Println(add(2, 5))
	
	//Consecutive typing
	fmt.Println(addCons(5, 2))
	
	//Multiple return values
	a, b := swapStrings("I", "say")
	fmt.Println(a, b)

	//Naked return function
	fmt.Println(naked(2, 2))

	//Func level vars
	//Can omit var with :=
	funcVarThree, funcVarFour := 1, 2;
	fmt.Println(packVarOne, packVarTwo)
	fmt.Println(funcVarThree, funcVarFour)

	//Var factorization
	var (
		h int = 255
		i string = "i var"
		j bool = true
	)
	fmt.Println("H: ", h)
	fmt.Println("I: ", i)
	fmt.Println("J: ", j)

	//Declared vars
	var (
		zerNum uint8
		zerStr string
		zerBool bool
	)
	fmt.Println(zerNum)
	fmt.Println(zerStr)
	fmt.Println(zerBool)

	//Casting
	smallVar := 42
	var largeVar int64
	largeVar = int64(smallVar)
	fmt.Println(smallVar, largeVar)
}

func exported() {
	fmt.Println(math.Pi)
}

// !Throws an exception
// Exported units begin with capital letter
//
// func exported() {
// 	  fmt.Println(math.pi)
// }

//Types come after arguments
func add(x int, y int) int {
	return x + y
}

//Consecutive arguments typing
func addCons(x, y int) int {
	return x + y
}

//Multiple return values
func swapStrings(first, second string) (string, string) {
	return second, first
}

//Return named values
func naked(firstVal, secondVal int) (x, y int) {
	x = firstVal * 2
	y = secondVal * 3
	return
}