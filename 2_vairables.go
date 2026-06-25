package main

// Importing the fmt package.
// fmt is used to print output to the console.
import "fmt"

// main() is the entry point of every Go program.
func main() {

	// ============================================================
	// 1. EXPLICIT VARIABLE DECLARATION
	// ============================================================

	// Syntax:
	// var variableName dataType = value

	// Here,
	// var    -> keyword used to declare a variable
	// name   -> variable name
	// string -> data type
	// "Go"   -> value assigned to the variable
	var name string = "Go"

	// Print the variable
	fmt.Println("Name:", name)

	// ============================================================
	// 2. TYPE INFERENCE
	// ============================================================

	// Go can automatically determine the data type from the value.
	// Since "Golang" is a string,
	// Go automatically makes language of type string.
	var language = "Golang"

	// Since 23 is an integer,
	// Go automatically makes age of type int.
	var age = 23

	fmt.Println("Language:", language)
	fmt.Println("Age:", age)

	// ============================================================
	// 3. SHORT VARIABLE DECLARATION (:=)
	// ============================================================

	// This is the most commonly used way to declare variables.
	// Go creates the variable, determines its type,
	// and assigns the value in one step.

	city := "Lucknow"

	fmt.Println("City:", city)

	// NOTE:
	// := can ONLY be used inside functions.
	// It cannot be used outside main() or other functions.

	// ============================================================
	// 4. DECLARE FIRST, ASSIGN LATER
	// ============================================================

	// Here we only declare the variable.
	// No value is assigned yet.
	var college string

	// Assigning value later.
	college = "IIIT Bhopal"

	fmt.Println("College:", college)

	// ============================================================
	// 5. INTEGER VARIABLES
	// ============================================================

	var marks int = 95

	// Type inference
	var semester = 8

	// Short declaration
	rollNumber := 101

	fmt.Println("Marks:", marks)
	fmt.Println("Semester:", semester)
	fmt.Println("Roll Number:", rollNumber)

	// ============================================================
	// 6. FLOAT VARIABLES
	// ============================================================

	// float32 occupies 32 bits of memory.
	var cgpa float32 = 8.25

	// Go automatically chooses float64 here.
	percentage := 81.75

	fmt.Println("CGPA:", cgpa)
	fmt.Println("Percentage:", percentage)

	// ============================================================
	// 7. BOOLEAN VARIABLES
	// ============================================================

	// Boolean variables can only store true or false.

	var isPlaced bool = false

	eligible := true

	fmt.Println("Placed:", isPlaced)
	fmt.Println("Eligible:", eligible)

	// ============================================================
	// 8. MULTIPLE VARIABLE DECLARATION
	// ============================================================

	// Multiple variables of the same type.
	var x, y, z int = 10, 20, 30

	fmt.Println(x, y, z)

	// Multiple variables with inferred types.
	var (
		company = "Google"
		salary  = 120000
	)

	fmt.Println(company)
	fmt.Println(salary)

	// ============================================================
	// 9. ZERO VALUES
	// ============================================================

	// If no value is assigned,
	// Go automatically assigns a default value.

	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Println("Default int:", defaultInt)         // 0
	fmt.Println("Default float:", defaultFloat)     // 0
	fmt.Println("Default string:", defaultString)   // ""
	fmt.Println("Default bool:", defaultBool)       // false

	// ============================================================
	// 10. CONSTANTS
	// ============================================================

	// Constants cannot be changed after initialization.

	const pi = 3.14159

	fmt.Println("PI:", pi)

	// Uncommenting the next line will produce a compilation error.
	// pi = 4.2

	// ============================================================
	// 11. VARIABLE REASSIGNMENT
	// ============================================================

	score := 80

	fmt.Println("Initial Score:", score)

	// Updating the value.
	score = 95

	fmt.Println("Updated Score:", score)

	// ============================================================
	// 12. UNUSED VARIABLES
	// ============================================================

	// Go does NOT allow unused local variables.
	// The following code would produce a compilation error.

	// temp := 100
	// Error:
	// declared and not used: temp

	// ============================================================
	// 13. INVALID REDECLARATION
	// ============================================================

	// Once a variable is declared,
	// you cannot declare it again in the same scope.

	// age := 25
	// Error:
	// no new variables on left side of :=

	// Instead, update its value.

	age = 24

	fmt.Println("Updated Age:", age)

	// ============================================================
	// 14. DIFFERENT DATA TYPES TOGETHER
	// ============================================================

	username := "Aditya"
	userAge := 23
	height := 177.5
	isStudent := false

	fmt.Println(username)
	fmt.Println(userAge)
	fmt.Println(height)
	fmt.Println(isStudent)

	// ============================================================
	// SUMMARY
	// ============================================================

	// Explicit declaration:
	// var name string = "Go"

	// Type inference:
	// var name = "Go"

	// Short declaration (most common):
	// name := "Go"

	// Declare first:
	// var name string
	// name = "Go"

	// Constant:
	// const pi = 3.14
}
