// package main
// As we saw before, we can bundle code in a function and call that function when we need the code inside it to run. We’ll be going over function syntax starting with a simple function definition:

// func summonNicole() {
//   fmt.Println("Hey Nicole, get over here!")
// }
// Above, we defined a function called summonNicole() and, within the body of the function (the part between the curly braces) we print out a message. It’s important to note that the code inside the function body does not run until we call the function. We call a function by using its name followed by parentheses somewhere outside the definition of the function. Our whole main.go file could look like this:

package main

import "fmt"

func summonNicole() {
	fmt.Println("Hey Nicole, get over here!")
}

func main() {
	// We call our function for the first time
	summonNicole()

	// We call our function for the second time
	summonNicole()
}

// In our example, we defined the function summonNicole() and called it twice inside our main() function. Notice that our function definition exists outside of main(), but calling summonNicole() occurs inside our main() function. This produces the following output in the terminal:

// Hey Nicole, get over here!
// Hey Nicole, get over here!
