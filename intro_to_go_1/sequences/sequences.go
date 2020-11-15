package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, j := range slice {
		slice[i] = f(j)
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for i, j := range array {
		array[i] = f(j)
	}
	fmt.Println("addOne to array", array)
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)
	fmt.Println("addOne to Slice", intsSlice)

	newSlice := intsSlice[1:3] // Given: intsSlice  = [2, 3, 4, 5, 6]
	mapSlice(square, newSlice) // newSlice = [3, 4]
	fmt.Println("square slice", newSlice)
	fmt.Println("square slice", intsSlice)

	intsSlice = double(intsSlice)
	fmt.Println("double slice", intsSlice)

	intsArray := [3]int{1, 2, 3}
	mapArray(addOne, intsArray)
	//fmt.Println(intsArray) //- doesn't work so print in mapArray function
	//arrays pass by value; slices pass by reference
}
