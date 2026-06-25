package main

import "fmt"

func main() {

	// ============================================================
	// 1. BASIC FOR LOOP
	// ============================================================

	// Syntax:
	// for initialization; condition; update {
	//     // code
	// }

	fmt.Println("1. Basic For Loop")

	for i := 1; i <= 5; i++ {
		fmt.Println("Value of i:", i)
	}

	// Output:
	// Value of i: 1
	// Value of i: 2
	// Value of i: 3
	// Value of i: 4
	// Value of i: 5

	// ============================================================
	// 2. PRINT EVEN NUMBERS
	// ============================================================

	fmt.Println("\n2. Even Numbers")

	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// ============================================================
	// 3. PRINT ODD NUMBERS
	// ============================================================

	fmt.Println("\n3. Odd Numbers")

	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// ============================================================
	// 4. COUNTDOWN
	// ============================================================

	fmt.Println("\n4. Countdown")

	for i := 5; i >= 1; i-- {
		fmt.Println(i)
	}

	// ============================================================
	// 5. SUM OF NUMBERS
	// ============================================================

	fmt.Println("\n5. Sum from 1 to 5")

	sum := 0

	for i := 1; i <= 5; i++ {
		sum += i
	}

	fmt.Println("Sum =", sum)

	// ============================================================
	// 6. WHILE LOOP IN GO
	// ============================================================

	// Go doesn't have a while keyword.
	// A for loop with only a condition behaves like a while loop.

	fmt.Println("\n6. While Loop")

	num := 1

	for num <= 5 {
		fmt.Println(num)
		num++
	}

	// ============================================================
	// 7. INFINITE LOOP
	// ============================================================

	// A for loop without any condition runs forever.

	// Uncomment to test.

	/*
		fmt.Println("\n7. Infinite Loop")

		for {
			fmt.Println("Running forever...")
		}
	*/

	// ============================================================
	// 8. BREAK
	// ============================================================

	fmt.Println("\n8. Break Statement")

	for i := 1; i <= 10; i++ {

		if i == 6 {
			break // Exit the loop
		}

		fmt.Println(i)
	}

	// Output:
	// 1 2 3 4 5

	// ============================================================
	// 9. CONTINUE
	// ============================================================

	fmt.Println("\n9. Continue Statement")

	for i := 1; i <= 5; i++ {

		if i == 3 {
			continue // Skip this iteration
		}

		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
	// 4
	// 5

	// ============================================================
	// 10. NESTED LOOPS
	// ============================================================

	fmt.Println("\n10. Nested Loop")

	for i := 1; i <= 3; i++ {

		for j := 1; j <= 3; j++ {

			fmt.Println("i =", i, "j =", j)

		}
	}

	// ============================================================
	// 11. LOOP OVER A SLICE (ARRAY-LIKE)
	// ============================================================

	fmt.Println("\n11. Range Loop")

	numbers := []int{10, 20, 30, 40}

	// range returns:
	// index and value

	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	// ============================================================
	// 12. IGNORE INDEX
	// ============================================================

	fmt.Println("\n12. Ignore Index")

	for _, value := range numbers {
		fmt.Println(value)
	}

	// "_" is called the blank identifier.
	// It ignores values you don't need.

	// ============================================================
	// 13. IGNORE VALUE
	// ============================================================

	fmt.Println("\n13. Ignore Value")

	for index := range numbers {
		fmt.Println(index)
	}

	// ============================================================
	// 14. LOOP OVER A STRING
	// ============================================================

	fmt.Println("\n14. String Loop")

	word := "GO"

	for index, ch := range word {

		fmt.Println(index, string(ch))

	}

	// ============================================================
	// 15. LOOP WITH MULTIPLE VARIABLES
	// ============================================================

	fmt.Println("\n15. Multiple Variables")

	for i, j := 1, 10; i <= 5; i, j = i+1, j-1 {

		fmt.Println(i, j)

	}

	// ============================================================
	// SUMMARY
	// ============================================================

	// Basic Loop
	// for i := 0; i < n; i++ {}

	// While Loop
	// for condition {}

	// Infinite Loop
	// for {}

	// Range Loop
	// for index, value := range collection {}

	// Break
	// break

	// Continue
	// continue
}
