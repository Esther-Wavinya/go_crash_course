// Projects can contain many .go files, organized into packages. Each package is like a directory: .go files to do with one part of your program can go in one package, other .go files to do with something else can go into another package. If we were writing a calculator program, we could put the files for the calculation in package calc and the files for input & output in package io.

// Package Declaration
// Let’s focus on the first line package main. This line is called a package declaration and it tells the Go compiler to which package this ‘.go’ file belongs. If this package declaration is ‘package main’, then this program will be compiled into an executable.

// Whitespace
// Next is a blank line. Go generally ignores these blank lines, they’re considered whitespace (new lines, spaces, and tabs). While our program doesn’t need the line break, it makes our code easier to read.

// Import Statement
// Then we have an import statement, import "fmt". The import keyword allows us to use code from other packages, in this case the Println function from the fmt package. Note how the package name is enclosed with double quotes ".

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// main Function
// We use the func keyword to declare the Go function main:

// The func keyword denotes the start of a function declaration.
// func is followed by the name of the function: main.
// After the name there follows a pair of parentheses () and a set of curly braces {}.
// The function code is written inside the set of curly braces.

// Function Code
// The code inside a function is indented. The code here invokes the Println function in the fmt package (that we imported earlier) to print the message "Hello World".

// Invoking Functions
// Normally when we write functions, you need to write code to invoke them, otherwise they are unused. However, the main function is different if it resides in the main package.

// Importing Multiple Packages
// Previously, we imported a single package, fmt. This package is useful but it is only one in a list of many included with Go.

// The standard packages are so useful that you will often use multiple packages in each .go file.

// Importing Multiple Packages
// To import multiple packages, we can add multiple import statements:

// import "package1"
// import "package2"
// Or use a single import with parentheses:

// import (
//   "package1"
//   "package2"
// )
// Package Aliases
// We can also provide an alias to a package by specifying an alias name before the package name.

// import (
//   p1 "package1"
//   "package2"
// )
// In the example above we’ve aliased package1 as p1 and now we can call functions from package1 by using p1 like:

// p1.SampleFunc()
// Instead of:

// package1.SampleFunc()
