package main

import "fmt"

// main() is the starting point of the program.
func main() {

	// ============================================================
	// 1. SIMPLE IF STATEMENT
	// ============================================================

	age := 20

	// If the condition is true,
	// the code inside the block executes.
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	}

	// ============================================================
	// 2. IF - ELSE
	// ============================================================

	marks := 35

	// If marks are greater than or equal to 40,
	// print Pass.
	// Otherwise, print Fail.

	if marks >= 40 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// ============================================================
	// 3. IF - ELSE IF - ELSE
	// ============================================================

	score := 82

	// Conditions are checked from top to bottom.
	// As soon as one condition is true,
	// the remaining conditions are skipped.

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 75 {
		fmt.Println("Grade B")
	} else if score >= 60 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	// ============================================================
	// 4. MULTIPLE CONDITIONS (&&)
	// ============================================================

	age = 22
	hasLicense := true

	// && means BOTH conditions must be true.

	if age >= 18 && hasLicense {
		fmt.Println("You can drive.")
	} else {
		fmt.Println("You cannot drive.")
	}

	// ============================================================
	// 5. USING OR (||)
	// ============================================================

	isWeekend := false
	isHoliday := true

	// || means ANY ONE condition should be true.

	if isWeekend || isHoliday {
		fmt.Println("Enjoy your holiday!")
	} else {
		fmt.Println("Go to work.")
	}

	// ============================================================
	// 6. USING NOT (!)
	// ============================================================

	isLoggedIn := false

	// ! reverses the boolean value.

	if !isLoggedIn {
		fmt.Println("Please login first.")
	}

	// ============================================================
	// 7. SHORT VARIABLE DECLARATION INSIDE IF
	// ============================================================

	// Variable 'num' exists only inside this if-else block.

	if num := 10; num%2 == 0 {
		fmt.Println(num, "is Even")
	} else {
		fmt.Println(num, "is Odd")
	}

	// ============================================================
	// 8. NESTED IF
	// ============================================================

	age = 25
	salary := 60000

	if age >= 18 {

		if salary >= 50000 {
			fmt.Println("Eligible for Premium Credit Card")
		} else {
			fmt.Println("Income is too low")
		}

	} else {
		fmt.Println("Age requirement not met")
	}

	// ============================================================
	// 9. COMPARISON OPERATORS
	// ============================================================

	a := 15
	b := 20

	if a == b {
		fmt.Println("a and b are equal")
	}

	if a != b {
		fmt.Println("a and b are NOT equal")
	}

	if a < b {
		fmt.Println("a is smaller")
	}

	if b > a {
		fmt.Println("b is greater")
	}

	if a <= b {
		fmt.Println("a is less than or equal to b")
	}

	if b >= a {
		fmt.Println("b is greater than or equal to a")
	}

	// ============================================================
	// 10. COMBINING CONDITIONS
	// ============================================================

	percentage := 88
	attendance := 80

	if percentage >= 75 && attendance >= 75 {
		fmt.Println("Eligible for Scholarship")
	} else {
		fmt.Println("Not Eligible")
	}

	// ============================================================
	// SUMMARY
	// ============================================================

	// Simple If
	// if condition {
	//     // code
	// }

	// If Else
	// if condition {
	//     // code
	// } else {
	//     // code
	// }

	// If Else If Else
	// if condition1 {
	// } else if condition2 {
	// } else {
	// }

	// Logical Operators
	// &&  -> AND
	// ||  -> OR
	// !   -> NOT
}
