package main

import "fmt"

func main() {

	// ============================================================
	// 1. RANGE WITH ARRAY
	// ============================================================

	fmt.Println("1. Range with Array")

	numbers := [5]int{10, 20, 30, 40, 50}

	// range returns:
	// index, value

	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	// ============================================================
	// 2. RANGE WITH SLICE
	// ============================================================

	fmt.Println("\n2. Range with Slice")

	fruits := []string{"Apple", "Mango", "Banana", "Orange"}

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	// ============================================================
	// 3. IGNORE INDEX
	// ============================================================

	fmt.Println("\n3. Ignore Index")

	// "_" is called the blank identifier.
	// It ignores values you don't need.

	for _, value := range fruits {
		fmt.Println(value)
	}

	// ============================================================
	// 4. IGNORE VALUE
	// ============================================================

	fmt.Println("\n4. Ignore Value")

	for index := range fruits {
		fmt.Println(index)
	}

	// ============================================================
	// 5. RANGE WITH MAP
	// ============================================================

	fmt.Println("\n5. Range with Map")

	student := map[string]int{
		"Math":    95,
		"Physics": 90,
		"Chem":    88,
	}

	// range returns:
	// key, value

	for subject, marks := range student {
		fmt.Println(subject, ":", marks)
	}

	// NOTE:
	// Maps are unordered.
	// Output order may change each time you run the program.

	// ============================================================
	// 6. RANGE WITH STRING
	// ============================================================

	fmt.Println("\n6. Range with String")

	word := "Golang"

	// range returns:
	// index and rune (Unicode character)

	for index, ch := range word {
		fmt.Println(index, string(ch))
	}

	// ============================================================
	// 7. RANGE WITH UTF-8 STRING
	// ============================================================

	fmt.Println("\n7. Unicode Example")

	text := "Go😊"

	for index, ch := range text {
		fmt.Println(index, string(ch))
	}

	// ============================================================
	// 8. MODIFY VALUES USING INDEX
	// ============================================================

	fmt.Println("\n8. Updating Slice")

	nums := []int{1, 2, 3, 4}

	// value is a COPY.
	// Use index to modify the original slice.

	for i := range nums {
		nums[i] *= 10
	}

	fmt.Println(nums)

	// ============================================================
	// 9. VALUE IS A COPY
	// ============================================================

	fmt.Println("\n9. Value is Copy")

	values := []int{5, 10, 15}

	for _, value := range values {
		value *= 2
	}

	// Original slice remains unchanged.

	fmt.Println(values)

	// ============================================================
	// 10. NORMAL FOR VS RANGE
	// ============================================================

	fmt.Println("\n10. Normal For")

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	fmt.Println("\nUsing Range")

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// ============================================================
	// SUMMARY
	// ============================================================

	// Array
	// for index, value := range arr {}

	// Slice
	// for index, value := range slice {}

	// Map
	// for key, value := range map {}

	// String
	// for index, char := range str {}

	// Ignore index
	// for _, value := range slice {}

	// Ignore value
	// for index := range slice {}
}
