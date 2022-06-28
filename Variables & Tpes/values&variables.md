# Values and Variables
Programs, like the ones we write in Go, are excellent at processing and performing operations on data. But in programming, “data” can be so many different things. Data can be numbers, boolean values (true/false), strings (blocks of text), or combinations thereof.

In this lesson, we’ll be covering how to store “data” by creating and using variables in Go.

We’re going to look into the different ways we can interact with these values (like adding two numbers together, or creating a longer message from two strings). We’re also going to discuss data types, the specific divisions that Go uses to “expect” certain qualities from variables. By creating and assigning values to our variables, we’ll understand the limitations and benefits of using different data types.

We’ll also cover how to read Go error messages — creating and using variables adds a new level of complexity to our programs to make these errors much more likely. This will be a good learning process, life as a programmer involves reading and interpreting these error messages quite often!


`main.go` like this: 

```
package main

import "fmt"

func main() {
  var stationName string
  var nextTrainTime int8
  var isUptownTrain bool
  
  stationName = "Union Square"
  nextTrainTime = 12
  isUptownTrain = false
  
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
  
  stationName = "Grand Central"
  nextTrainTime = 3
  isUptownTrain = true
  
  fmt.Println("")
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
}
```

`Output` after running `main.go`

```
Current station: Union Square
Next train: 12 minutes
Is uptown: false

Current station: Grand Central
Next train: 3 minutes
Is uptown: true
```


# Literals
In Go, values can be many things. Just to name a few, values can be numbers (like `109`), or text wrapped in quotes (like `"Hello world"`). These values can be written into code as is, and are called *literals*. They are literally what they say they are.

We can perform arithmetic in Go with literals (or named values, covered in the next exercise) using the following operators:

- `+` to add
- `-` to subtract
- `*` to multiply
- `/` to divide
- `%` to take the remainder (the modulus operator) between two numbers.

```
fmt.Println(20 * 3) // Prints: 60
fmt.Println(55.21 / 5) // Prints: 11.042
fmt.Println(9 % 2) // Prints: 1
```

Imagine the code above as appearing inside the `main` function body of a Go program. In this snippet, we used the Go programming language as a calculator. We printed out the product of multiplying `20` and `3` (`60`). We next printed out the quotient of dividing `55.21` by `5` (`11.042`). We lastly printed out the remainder left over after dividing `9` by `2` (9 divided by `2` has a remainder of `1`).

`main.go` like this:

```
package main

import "fmt"

func main() {
  // Add a fmt.Println() statement
  // that prints 2235 * 1231
  fmt.Println(2235 * 1231)
}
```

`Output` after running `main.go` is: 

```
2751285
```


