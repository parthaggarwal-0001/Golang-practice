package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}

	mySlice := arr1[2:4]

	fmt.Printf("Slice %v\n", mySlice)
	fmt.Printf("Array %v\n", arr1)

	fmt.Printf("Length %d\n", len(mySlice))
	fmt.Printf("Capacity %d\n", cap(mySlice))

	mySlice = append(mySlice, 2, 3)

	fmt.Printf("After append Slice %v\n", mySlice)
	fmt.Printf("After append Array %v\n", arr1)
	
}
