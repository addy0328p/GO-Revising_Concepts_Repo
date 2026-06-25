package main

import "fmt"

// main() is the entry point of the program.
func main() {

	// ============================================================
	// 1. WHAT IS A SLICE?
	// ============================================================

	// A slice is a dynamic (resizable) view of an array.
	// Unlike arrays, slices do NOT have a fixed size.

	// Syntax:
	// var sliceName []dataType

	var numbers []int

	// Currently the slice is empty.
	fmt.Println("Empty Slice:", numbers)

	// Length = number of elements
	fmt.Println("Length:", len(numbers))

	// Capacity = maximum elements before Go allocates new memory
	fmt.Println("Capacity:", cap(numbers))

	// ============================================================
	// 2. DECLARING A SLICE WITH VALUES
	// ============================================================

	fruits := []string{"Apple", "Mango", "Banana", "Orange"}

	fmt.Println("\nFruits:", fruits)

	// ============================================================
	// 3. ACCESSING ELEMENTS
	// ============================================================

	fmt.Println("First Fruit:", fruits[0])
	fmt.Println("Second Fruit:", fruits[1])
	fmt.Println("Last Fruit:", fruits[3])

	// ============================================================
	// 4. MODIFYING ELEMENTS
	// ============================================================

	fruits[1] = "Pineapple"

	fmt.Println("\nUpdated Fruits:", fruits)

	// ============================================================
	// 5. APPENDING ELEMENTS
	// ============================================================

	// append() adds new elements to a slice.

	fruits = append(fruits, "Grapes")

	fmt.Println("\nAfter Append:", fruits)

	// Add multiple values

	fruits = append(fruits, "Kiwi", "Guava")

	fmt.Println("After Multiple Append:", fruits)

	// ============================================================
	// 6. LENGTH AND CAPACITY
	// ============================================================

	fmt.Println("\nLength:", len(fruits))
	fmt.Println("Capacity:", cap(fruits))

	// ============================================================
	// 7. LOOPING USING FOR
	// ============================================================

	fmt.Println("\nNormal For Loop")

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// ============================================================
	// 8. LOOPING USING RANGE
	// ============================================================

	fmt.Println("\nRange Loop")

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	// ============================================================
	// 9. IGNORE INDEX
	// ============================================================

	fmt.Println("\nIgnore Index")

	for _, value := range fruits {
		fmt.Println(value)
	}

	// ============================================================
	// 10. IGNORE VALUE
	// ============================================================

	fmt.Println("\nIgnore Value")

	for index := range fruits {
		fmt.Println(index)
	}

	// ============================================================
	// 11. CREATING A SUB-SLICE
	// ============================================================

	// Syntax:
	// slice[start:end]
	// Includes start
	// Excludes end

	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("\nOriginal:", nums)

	fmt.Println("nums[1:4] =", nums[1:4]) // 20 30 40
	fmt.Println("nums[:3] =", nums[:3])   // 10 20 30
	fmt.Println("nums[2:] =", nums[2:])   // 30 40 50

	// ============================================================
	// 12. MAKE FUNCTION
	// ============================================================

	// make(type, length, capacity)

	scores := make([]int, 5)

	fmt.Println("\nUsing make():", scores)

	fmt.Println("Length:", len(scores))
	fmt.Println("Capacity:", cap(scores))

	// Assign values

	scores[0] = 90
	scores[1] = 85
	scores[2] = 78

	fmt.Println(scores)

	// ============================================================
	// 13. MAKE WITH CAPACITY
	// ============================================================

	data := make([]int, 3, 10)

	fmt.Println("\nSlice:", data)
	fmt.Println("Length:", len(data))
	fmt.Println("Capacity:", cap(data))

	// ============================================================
	// 14. COPYING SLICES
	// ============================================================

	source := []int{1, 2, 3, 4}

	destination := make([]int, len(source))

	copy(destination, source)

	fmt.Println("\nSource:", source)
	fmt.Println("Destination:", destination)

	// ============================================================
	// 15. NIL SLICE
	// ============================================================

	var s []int

	fmt.Println("\nNil Slice:", s)

	if s == nil {
		fmt.Println("Slice is nil")
	}

	// ============================================================
	// 16. REMOVE ELEMENT
	// ============================================================

	values := []int{10, 20, 30, 40, 50}

	// Remove 30 (index 2)

	values = append(values[:2], values[3:]...)

	fmt.Println("\nAfter Removing 30:", values)

	// ============================================================
	// SUMMARY
	// ============================================================

	// Declare
	// var s []int

	// Declare with values
	// s := []int{1,2,3}

	// Create using make
	// make([]int, length, capacity)

	// Add elements
	// append(slice, value)

	// Length
	// len(slice)

	// Capacity
	// cap(slice)

	// Copy
	// copy(destination, source)

	// Range
	// for index, value := range slice {}
}
