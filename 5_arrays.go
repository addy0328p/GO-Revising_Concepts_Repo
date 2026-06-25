package main

import "fmt"

// main() is the entry point of the program.
func main() {

	// ============================================================
	// 1. WHAT IS AN ARRAY?
	// ============================================================

	// An array is a fixed-size collection of elements
	// of the SAME data type.

	// Syntax:
	// var arrayName [size]dataType

	var numbers [5]int

	// At this point all elements contain the zero value.
	// int -> 0

	fmt.Println("Empty Array:", numbers)

	// ============================================================
	// 2. ASSIGNING VALUES
	// ============================================================

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Numbers:", numbers)

	// ============================================================
	// 3. ACCESSING ELEMENTS
	// ============================================================

	fmt.Println("First Element :", numbers[0])
	fmt.Println("Third Element :", numbers[2])
	fmt.Println("Last Element  :", numbers[4])

	// ============================================================
	// 4. ARRAY DECLARATION WITH VALUES
	// ============================================================

	// Values can be assigned during declaration.

	fruits := [4]string{"Apple", "Mango", "Banana", "Orange"}

	fmt.Println("Fruits:", fruits)

	// ============================================================
	// 5. ARRAY LENGTH
	// ============================================================

	// len() returns the number of elements.

	fmt.Println("Length of fruits:", len(fruits))

	// ============================================================
	// 6. LOOPING THROUGH AN ARRAY
	// ============================================================

	fmt.Println("\nUsing normal for loop")

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// ============================================================
	// 7. USING range
	// ============================================================

	fmt.Println("\nUsing range")

	// range returns:
	// index, value

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	// ============================================================
	// 8. IGNORING INDEX
	// ============================================================

	fmt.Println("\nIgnoring Index")

	for _, value := range fruits {
		fmt.Println(value)
	}

	// ============================================================
	// 9. IGNORING VALUE
	// ============================================================

	fmt.Println("\nIgnoring Value")

	for index := range fruits {
		fmt.Println(index)
	}

	// ============================================================
	// 10. ARRAY OF FLOATS
	// ============================================================

	prices := [3]float64{99.5, 120.75, 300.20}

	fmt.Println("\nPrices:", prices)

	// ============================================================
	// 11. ARRAY OF BOOLEANS
	// ============================================================

	flags := [4]bool{true, false, true, false}

	fmt.Println("Flags:", flags)

	// ============================================================
	// 12. AUTOMATIC SIZE (...)
	// ============================================================

	// Go counts the number of elements automatically.

	cities := [...]string{
		"Lucknow",
		"Delhi",
		"Mumbai",
		"Pune",
	}

	fmt.Println("\nCities:", cities)
	fmt.Println("Length:", len(cities))

	// ============================================================
	// 13. COPYING ARRAYS
	// ============================================================

	a := [3]int{1, 2, 3}

	// Arrays are copied by value.
	b := a

	// Changing b does NOT affect a.
	b[0] = 100

	fmt.Println("\nOriginal Array :", a)
	fmt.Println("Copied Array   :", b)

	// ============================================================
	// 14. COMPARING ARRAYS
	// ============================================================

	x := [3]int{1, 2, 3}
	y := [3]int{1, 2, 3}

	if x == y {
		fmt.Println("\nArrays are equal")
	}

	// ============================================================
	// 15. MULTI-DIMENSIONAL ARRAY
	// ============================================================

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("\nMatrix:", matrix)

	fmt.Println("matrix[0][0] =", matrix[0][0])
	fmt.Println("matrix[1][2] =", matrix[1][2])

	// ============================================================
	// 16. MODIFYING AN ELEMENT
	// ============================================================

	numbers[2] = 999

	fmt.Println("\nUpdated Numbers:", numbers)

	// ============================================================
	// 17. ZERO VALUES
	// ============================================================

	var marks [4]int
	var names [3]string
	var passed [2]bool

	fmt.Println("\nDefault int array    :", marks)
	fmt.Println("Default string array :", names)
	fmt.Println("Default bool array   :", passed)

	// ============================================================
	// SUMMARY
	// ============================================================

	// Declare
	// var arr [5]int

	// Declare with values
	// arr := [5]int{1,2,3,4,5}

	// Automatic size
	// arr := [...]int{1,2,3}

	// Access
	// arr[0]

	// Update
	// arr[1] = 100

	// Length
	// len(arr)

	// Loop
	// for i := 0; i < len(arr); i++ {}

	// Range
	// for index, value := range arr {}
}
