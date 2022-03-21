// Instructions
// In this exercise you're going to write some code to help you cook a brilliant lasagna from your favorite cooking book.

// You have four tasks, all related to the time spent cooking the lasagna.


// 1. Define the expected oven time in minutes
// Define the OvenTime constant with how many minutes the lasagna should be in the oven. According to the cooking book, the expected oven time in minutes is 40:

// OvenTime
// // Output: 40
// 
// Define the expected oven time in minutes
// You need to define a constant and assign it the expected oven time in minutes.

// If you see an undefined: OvenTime error then double check that you have the constant defined.

// If you see an invalid operation: got != tt.expected (mismatched types float64 and int) error then you have likely put a decimal point into the OvenTime causing Go to infer the type as a floating point number. Remove the decimal and the type will be inferred as an int.

// If you see a syntax error: non-declaration statement outside function body error then it is likely that you forgot the const keyword.

// If you see a syntax error: unexpected :=, expecting = error then you are likely trying to assign the constant using := like a variable; constants are assigned using = not :=.

// TODO: define the 'OvenTime' constant
package lasagna

const OvenTime = 40





// 2. Calculate the remaining oven time in minutes
// Define the RemainingOvenTime() function that takes the actual minutes the lasagna has been in the oven as a parameter and returns how many minutes the lasagna still has to remain in the oven, based on the expected oven time in minutes from the previous task.

// func RemainingOvenTime(actual int) int {
//     // TODO
// }

// RemainingOvenTime(30)
// // Output: 10
// Calculate the remaining oven time in minutes
// You need to define a function with a single parameter.

// You have to explicitly return an integer from a function.

// The function's parameter is an integer.

// You can call one of the other functions you've defined previously.

// You can use the mathematical operator for subtraction to subtract values.

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    return(OvenTime - actualMinutesInOven)	
}





// 3. Calculate the preparation time in minutes
// Define the PreparationTime function that takes the number of layers you added to the lasagna as a parameter and returns how many minutes you spent preparing the lasagna, assuming each layer takes you 2 minutes to prepare.

// func PreparationTime(numberOfLayers int) int {
//     // TODO
// }

// PreparationTime(2)
// // // Output: 4
// Calculate the preparation time in minutes
// You need to define a function with a single parameter.

// You have to explicitly return an integer from a function.

// The function's parameter is an integer.

// You can use the mathematical operator for multiplication to multiply values.


// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    return(numberOfLayers * 2)	
}





// 4. Calculate the elapsed working time in minutes
// Define the ElapsedTime function that takes two parameters: the first parameter is the number of layers you added to the lasagna, and the second parameter is the number of minutes the lasagna has been in the oven. The function should return how many minutes in total you've worked on cooking the lasagna, which is the sum of the preparation time in minutes, and the time in minutes the lasagna has spent in the oven at the moment.

// func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
//     // TODO
// }

// ElapsedTime(3, 20)
// // Output: 26

// Calculate the elapsed working time in minutes
// You need to define a function with two parameters.

// You have to explicitly return an integer from a function.

// The function's parameter is an integer.

// You can call one of the other functions you've defined previously.

// You can use the mathematical operator for addition to add values.

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return(PreparationTime(numberOfLayers) + actualMinutesInOven)
 }
 








//  Introduction
// Go is a statically typed, compiled programming language. This exercise introduces three major language features: Packages, Functions, and Variables.

// Packages
// Go applications are organized in packages. A package is a collection of source files located in the same directory. All source files in a directory must share the same package name. When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter can be used / accessed. The recommended style of naming in Go is that identifiers will be named using camelCase, except for those meant to be accessible across packages which should be PascalCase.

// package lasagna
// Variables
// Go is statically-typed, which means all variables must have a defined type at compile-time.

// Variables can be defined by explicitly specifying a type:

// var explicit int // Explicitly typed
// You can also use an initializer, and the compiler will assign the variable type to match the type of the initializer.

// implicit := 10   // Implicitly typed as an int
// Once declared, variables can be assigned values using the = operator. Once declared, a variable's type can never change.

// count := 1 // Assign initial value
// count = 2  // Update to new value

// count = false // This throws a compiler error due to assigning a non `int` type
// Constants
// Constants hold a piece of data just like variables, but their value cannot change during the execution of the program.

// Constants are defined using the const keyword and can be numbers, characters, strings or booleans:

// const Age = 21 // Defines a numeric constant 'Age' with the value of 21
// Functions
// Go functions accept zero or more parameters. Parameters must be explicitly typed, there is no type inference.

// Values are returned from functions using the return keyword.

// A function is invoked by specifying the function name and passing arguments for each of the function's parameters.

// Note that Go supports two types of comments. Single line comments are preceded by // and multiline comments are inserted between /* and */.

// package greeting

// // Hello is a public function
// func Hello (name string) string {
//     return hi(name)
// }

// // hi is a private function
// func hi (name string) string {
//     return "hi " + name
// }