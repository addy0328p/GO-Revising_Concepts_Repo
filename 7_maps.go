package main

import "fmt"

// main() is the entry point of the program.
func main() {

	// ============================================================
	// 1. WHAT IS A MAP?
	// ============================================================

	// A map stores data in key-value pairs.
	//
	// Syntax:
	// var mapName map[keyType]valueType

	var student map[string]string

	// Currently this map is nil.
	fmt.Println("Nil Map:", student)

	// NOTE:
	// You CANNOT insert values into a nil map.
	// student["Name"] = "Aditya" // ❌ Runtime panic

	// ============================================================
	// 2. CREATING A MAP USING make()
	// ============================================================

	// make() allocates memory for the map.

	student = make(map[string]string)

	// Adding key-value pairs

	student["Name"] = "Aditya"
	student["College"] = "IIIT Bhopal"
	student["City"] = "Lucknow"

	fmt.Println("\nStudent Map:", student)

	// ============================================================
	// 3. ACCESSING VALUES
	// ============================================================

	fmt.Println("\nName:", student["Name"])
	fmt.Println("City:", student["City"])

	// ============================================================
	// 4. UPDATING VALUES
	// ============================================================

	student["City"] = "Delhi"

	fmt.Println("\nUpdated City:", student["City"])

	// ============================================================
	// 5. DELETING A KEY
	// ============================================================

	delete(student, "College")

	fmt.Println("\nAfter Delete:", student)

	// ============================================================
	// 6. CHECK IF A KEY EXISTS
	// ============================================================

	value, exists := student["Name"]

	if exists {
		fmt.Println("\nKey Found:", value)
	} else {
		fmt.Println("Key Not Found")
	}

	// Checking a missing key

	_, exists = student["Phone"]

	if !exists {
		fmt.Println("Phone key doesn't exist")
	}

	// ============================================================
	// 7. MAP LITERAL
	// ============================================================

	employee := map[string]int{
		"John":  50000,
		"Alice": 70000,
		"Bob":   65000,
	}

	fmt.Println("\nEmployee Map:", employee)

	// ============================================================
	// 8. LOOPING THROUGH A MAP
	// ============================================================

	fmt.Println("\nUsing range")

	for key, value := range employee {
		fmt.Println(key, "->", value)
	}

	// NOTE:
	// Maps are unordered.
	// The output order may change every time you run the program.

	// ============================================================
	// 9. IGNORE KEY
	// ============================================================

	fmt.Println("\nValues Only")

	for _, value := range employee {
		fmt.Println(value)
	}

	// ============================================================
	// 10. IGNORE VALUE
	// ============================================================

	fmt.Println("\nKeys Only")

	for key := range employee {
		fmt.Println(key)
	}

	// ============================================================
	// 11. LENGTH OF A MAP
	// ============================================================

	fmt.Println("\nNumber of Entries:", len(employee))

	// ============================================================
	// 12. MAP WITH DIFFERENT VALUE TYPES
	// ============================================================

	marks := map[string]float64{
		"Math":    95.5,
		"Physics": 91.0,
		"Chem":    88.5,
	}

	fmt.Println("\nMarks:", marks)

	// ============================================================
	// 13. MAP OF BOOLEAN VALUES
	// ============================================================

	isLoggedIn := map[string]bool{
		"Aditya": true,
		"Rahul":  false,
	}

	fmt.Println("\nLogin Status:", isLoggedIn)

	// ============================================================
	// 14. NESTED MAP
	// ============================================================

	users := map[string]map[string]string{

		"user1": {
			"Name": "Aditya",
			"City": "Lucknow",
		},

		"user2": {
			"Name": "Rahul",
			"City": "Delhi",
		},
	}

	fmt.Println("\nNested Map:")
	fmt.Println(users)

	// Access nested values

	fmt.Println(users["user1"]["Name"])

	// ============================================================
	// 15. ZERO VALUE
	// ============================================================

	// If a key does not exist,
	// Go returns the zero value of the value type.

	ages := map[string]int{
		"Aditya": 23,
	}

	fmt.Println("\nAge:", ages["Aditya"])

	// "Rahul" doesn't exist.
	// int zero value = 0

	fmt.Println("Unknown:", ages["Rahul"])

	// ============================================================
	// SUMMARY
	// ============================================================

	// Declare
	// var m map[string]int

	// Initialize
	// m = make(map[string]int)

	// Literal
	// m := map[string]int{
	//     "A": 1,
	// }

	// Add
	// m["A"] = 10

	// Read
	// m["A"]

	// Update
	// m["A"] = 20

	// Delete
	// delete(m, "A")

	// Check existence
	// value, ok := m["A"]

	// Length
	// len(m)

	// Loop
	// for key, value := range m {}
}
