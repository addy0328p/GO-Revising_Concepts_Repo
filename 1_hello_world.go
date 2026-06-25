// Every Go program starts with a package declaration.
// 'main' is a special package that tells Go this is an executable program.
package main

// Import is used to include packages (libraries) that provide extra functionality.
// 'fmt' is Go's standard formatting package.
// We use it for printing text to the terminal.
import "fmt"

// main() is the entry point of every executable Go program.
// When you run the program, execution starts from this function.
func main() {

	// fmt.Println() prints the given text to the console.
	// 'Println' means "Print Line":
	// - Prints the text.
	// - Automatically moves the cursor to the next line.
	fmt.Println("hello world")
}
