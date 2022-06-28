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


# Constants
In addition to literal (aka *unnamed*) values, there are also named values. Naming a value in Go means creating a word that will represent that value. One example of named values are *constants*, which cannot be updated while the program is running. Another example of named values are *variables* which we can update their value. We’ll focus on constants in this exercise and discuss variables at length in later exercises.

We use the `const` keyword to create a constant. We immediately assign a value to the constant using a `literal`. Throughout the rest of the program, we can use the constant’s name instead of the literal.

```
const funFact = "Hummingbirds' wings can beat up to 200 times a second."
 
fmt.Println("Did you know?")
fmt.Println(funFact)
```

Which prints:

```
Did you know? 
Hummingbirds' wings can beat up to 200 times a second.
```

Above, we created a constant named `funFact`, which contains the text `"Hummingbirds' wings can beat up to 200 times a second."`. We were then able to print out this value using the names we applied. This is useful if a value doesn’t change throughout a program and it also helps to convey the developer’s intention of keeping a consistent value.

Let’s also take a second to examine the peculiar name, `funFact`. Constants have names without spaces: spaces aren’t allowed! This means that for us to have descriptive names (and it is important to have descriptive names so that we can read the code we’ve written) we need to use camelCase or PascalCase, capitalizing each subsequent word instead of adding spaces.

`main.go` like this:

```
package main

import "fmt"

func main() {
  // Create the constant earthsGravity
  // here and set it to 9.80665
  const earthsGravity = 9.80665
  // Here's where we print out the gravity:
  fmt.Println(earthsGravity)
}
```

`Outuput` after running `main.go` is:

```
9.80665
```



# What is a Data Type?
Programming languages need to process and organize data. That data is stored as binary numbers (numbers consisting of 0’s and 1’s) in the memory of your computer. In this way, binary numbers are used to represent many different things. *Data types* are designations by a programming language about what information is being stored. Go comes with a number of basic types, data types built into the programming language. It’s also possible to create your own types that combine these basic types into something more complex, but let’s first cover the data that you can store by default in Go.

Go has three basic categories for numbers:

- Integers, or `int`s, are whole/counting numbers. You would use an `int` to count the number of books on a shelf, the number of products in a warehouse, the number of people on a website, etc…

- Floating-point numbers, or `float`s, can include fractional data. You would use a float to store distances, percentages, and other quantities that required division or precision.

- Complex numbers, `complex`, are pairs of floating-point numbers where the second part of the pair is marked with the “imaginary” unit `i`. Complex numbers are particularly useful when reasoning in 2-dimensional space and have other utilizations that make them relevant for involved calculations, but we won’t be discussing them at length in this course.

![Go Basic Data Types](../assets/basic%20datatypes%20go.jpg)


# Basic Numeric Types in Go
Go has 15 different numeric types that fall into the three categories: `int`, `float`, and `complex`. That means there are fifteen different ways to describe a number in Go. This includes 11 different integer types, 2 different floating-point types, and 2 different complex number types. These types all recognize different sets of numbers as valid. An integer can’t store the number `8.6132`, for instance.

Beyond being broken down into the three categories, types also indicate how many bits (binary digits) will be used to represent the data. Fewer bits means fewer different possible values for the data, enforced as a strict minimum and maximum value for integers and less precision for floats and complex numbers. Fewer bits also means less data to save, so it will use less of a computer’s memory to hold onto that data. So, while it may be tempting to use types that can take a larger range of values, it can slow down a computer’s performance or cause the computer to run out of memory.

Integers are further broken down into two categories: `signed` and `unsigned`. Signed integers can be negative, but unsigned integers can only be positive. This means that the minimum value for an unsigned integer is always 0. Since it can ignore the possibility of a negative value, an unsigned integer’s maximum value is much higher than a signed integer with the same number of bits.

Go also has a *boolean* type. Booleans are either `false` or `true`. Go only needs one bit to store a boolean value: `0` represents `false` and `1` represents `true`.

![basic numeric types in Go](../assets/basic%20numeric%20types.jpg)

