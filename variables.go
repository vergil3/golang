// variables & constants
package main

import f "fmt"

func main() {
	f.Println("This is a program about variables and constants")

	//a constant is a variable with a value that can't be changed
	const pi = 3.14159

	//declare multiple variables like this
	var (
		variableA int = 2
		variableB int = 3
	)

	f.Println(variableA, variableB)

}
