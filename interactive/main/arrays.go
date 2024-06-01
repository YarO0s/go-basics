package main

import "fmt"

func main() {
	arrays()
	slices()
	viewSlices()
	sliceLiterals()
	makeArray()
	_ = twoDimArray()
	changeSize()
	iterate()
}

func arrays() {
	//array declaration
	var fixedArr [3]int
	fixedArr = [3]int{0,1,2}
	fmt.Println(fixedArr)

	infFixedArr := [5]int{0,1,2,3,4}
	fmt.Println(infFixedArr)

	value := fixedArr[0]
	fmt.Println(value)
}

func slices() {
	nums := [10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(nums[1:4])
}

func viewSlices() {
	arr := [10]int{0,1,2,3,4,5,6,7,8,9}
	firstSlice := arr[1:5]
	secondSlice := arr[4:7]
	
	fmt.Println(
		"Intersect in array: ", arr[4], 
		"; firstSlice: ", firstSlice[3], 
		"; secondSlice: ", secondSlice[0])

	secondSlice[0] *= 2
	
	fmt.Println(
		"After secondSlice[0] modified, array: ", arr[4],
		"; firstSlice: ", firstSlice[3],
		"; secondSlice: ", secondSlice[0])
}

// []type{val1, val2} slice literal with dynamic lenght
func sliceLiterals() {
	a := []int{0,1,2,3,4,5}
	fmt.Println(a)

	digits := []struct {
		a int
		b string
	} {
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
	}

	fmt.Println(digits)
}

func makeArray() {
	s := make([]int, 0, 10)
	fmt.Println(s)
}

func twoDimArray() ([][]string) {
	arr := [][]string {
		[]string {"zero", "one", "two"},
		[]string {"three", "four", "five"},
		[]string {"six", "seven", "eight"},
		[]string {"nine"},
	}

	fmt.Println(arr[0][0])
	return arr
}

//append helps reshape original slice (and backed array if needed)
func changeSize() {
	arr := []int {5}
	extendedSlice := append(arr, 25, 125)
	fmt.Println(extendedSlice)
}

//range as a foreach
func iterate() {
	arr := []string {"Hello ", "my ", "friend."}

	for _, val := range arr {
		fmt.Print(val)
	}
}