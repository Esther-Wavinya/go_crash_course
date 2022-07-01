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


# What is a Variable?
Now that we have some background on what types are, we can talk about what variables are and how we make and use them. A variable is a named value (like a constant) with the added feature that it can change during the running of a program. If we’re working with a value in various places in our program, we can store that value in a variable to easily access it later.

Variables are defined with the `var` keyword and two pieces of information: the name that we will use to refer to them and the type of data stored in the variable. Since variables can be updated we don’t even need to assign a value initially. Here’s a couple of variable definitions:

```
var lengthOfSong uint16
var isMusicOver bool
var songRating float32
```
  
Above, we created three variables:

- An unsigned 16-bit integer called `lengthOfSong`.
- A boolean value called `isMusicOver`.
- A 32-bit float called `songRating`.
  
Notice that our variable names also follow the same naming convention of constants, using camelCase with a descriptive name.

Create a variable called jellybeanCounter, that will store the number of jellybeans in a jar. Give it a type of int8.

Successfully doing this will result in an error! We’ll talk more about errors in Go and how to read them in the next exercise, so continue when you’re ready.

`main.go` like this: 

```
package main

func main() {
  // Create the variable jellybeanCounter
  // here and make its type int8
  var jellybeanCounter int8
  
}
```

`Output` will be:

```
# command-line-arguments
./main.go:6:7: jellybeanCounter declared but not used
```


# Reading Go Errors
There is no shame in having your code fail to run. Programming errors and exceptions happen all the time and learning to read and understand them is an indispensable tool in a programmer’s toolbox.

When the Go compiler raises an error the program’s binary cannot get created and without the binary, our computer cannot execute the program’s code. Go tries to tell you what the issue is by offering you the following pieces of information: the name of the file where something went wrong, the line number in that file where Go noticed an issue, and the column number (the number of characters from the left) on that line where the error occurred. One common error occurs when Go language’s compiler recognizes that there is a defined variable that isn’t used in the program. Take for example:

```
package main
 
func main() {
  var numberWheels int8
}
```

When we attempt to run `main.go` and on line 4 of our file we define the variable `numberWheels` that we don’t use anywhere else in the program we will see the following message:

```
./main.go:4:7: numberWheels declared and not used
```

Notice that the error message contains the file name (`main.go`), the line that causes the error to be raised (`line 4`) and specifically the location (`7` characters to the right of the line break). If we remove the variable from our program or use it somewhere we appease Go’s compiler and can run our program.

This may seem like a downside to a language, something that keeps you from expressing your freedom and personality as a programmer, but it is designed to discover possible errors in your program sooner rather than later. An unused variable is a waste of memory, but it can also be a typo or an unintended omission. In this case, Go’s somewhat hardline stance is to keep programmers from wondering why their code isn’t working the way they expect by refusing to run until some action is taken on this unused definition.

`main.go`

```
package main

import "fmt"

func main() {
  var jellybeanCounter int8
  
  // Go will raise errors both for
  // jellybeanCounter being unused
  // and for "fmt" being imported
  // and unused.
  
  // Uncomment the following line
  // and watch the program run
  // successfully!
  
  //fmt.Println(jellybeanCounter)
}
```

`Output`
```# command-line-arguments
./main.go:3:8: imported and not used: "fmt"
```

Run the code and review the error message. This time it’s complaining about “fmt”, a package we imported in order to print out the contents of jellybeanCounter. We’ll solve both this error and our error from the previous exercise by using both in the line commented at the bottom of main.go. Uncomment the line by removing the // at the beginning, and watch the program run!

`main.go` after uncommenting

```
package main

import "fmt"

func main() {
  var jellybeanCounter int8
  
  // Go will raise errors both for
  // jellybeanCounter being unused
  // and for "fmt" being imported
  // and unused.
  
  // Uncomment the following line
  // and watch the program run
  // successfully!
  
  fmt.Println(jellybeanCounter)
}
```

`Output` after uncommenting is:

```
0
```


# Assigning Variables
Variables are placeholder names that we use to refer to values that we intend to update over the course of our program. Updating our variable is also called *assigning* a value to a variable. In order to assign values to variables, we can use the assignment operator (`=`) followed by a value.

```
var kilometersToMars int32
 
kilometersToMars = 62100000
```

In the example above, we first declare our variable using the `var` keyword, the name `kilometersToMars`, and its type `int32`. Then, we assign `62100000` to `kilometersToMars`. Now when we need to know how many kilometers it is to mars, we can access the value using `kilometersToMars`!

Another way to assign our variable is:

```
var kilometersToMars int32 = 62100000
```

In our latest example, we declare `kilometersToMars`, assign the type (`int32`) and *initialize* it (assign a variable it’s first value) with a value of `62100000`. This syntax is helpful if we know exactly what type we want our variable to have and if we know what specific value to initialize it to.

- In `main.go`, we’ve declared a `int` variable named numOfFlavors.
Assign a value of `57` to `numOfFlavors`.
```
package main

import "fmt"

func main() {
  var numOfFlavors int
  // Assign a value for numOfFlavors below:
  numOfFlavors = 57
  
  fmt.Println(numOfFlavors)
  
  // Declare flavorScale below:
  
  
}
```

- Under the `print` statement, declare a variable named `flavorScale` that has a type of `float32` and initialize it with a value of `5.8` all on a single line.

```
package main

import "fmt"

func main() {
  var numOfFlavors int
  // Assign a value for numOfFlavors below:
  numOfFlavors = 57
  
  fmt.Println(numOfFlavors)
  
  // Declare flavorScale below:
  var flavorScale float32 = 5.8
  
}
```

- To check that our flavorScale variable was correctly defined, print out `flavorScale‘s value` using `fmt.Println()`.

```
package main

import "fmt"

func main() {
  var numOfFlavors int
  // Assign a value for numOfFlavors below:
  numOfFlavors = 57
  
  fmt.Println(numOfFlavors)
  
  // Declare flavorScale below:
  var flavorScale float32 = 5.8
  
  fmt.Println(flavorScale)
}
```

# Strings
We’ve talked about numeric types so far, but Go offers a few other built-in types. One particularly useful type is called a *string*. A string is Go’s type for storing and processing text. In a general programming sense, “string” is a term for text of any length, and the name is chosen to evoke a *sequence* of data.

Below, we declared two string variables:

```
var nameOfSong string
var nameOfArtist string
```

Afterward, we assign a value to the variables with the `=` sign. Surround the text you want to store with double-quotation marks (i.e., `"`, the single-quote `'` has a special other meaning and isn’t used to define strings):

```
nameOfSong = "Stop Stop"
nameOfArtist = "The Julie Ruin"
```

You can use the `+` operator on strings to join them, this is known as *string concatenation*. Note that `+` does not add in additional spaces or punctuation.

```
var description string
description = nameOfSong + " is by: " + nameOfArtist + "."
fmt.Println(description)
// Prints "Stop Stop is by: The Julie Ruin.
```

- In `main.go`, create a string variable called `favoriteSnack` but do not assign a value just yet.
```
package main

import "fmt"

func main() {
  // Define a string variable
  // called favoriteSnack

  var favoriteSnack string

  // Assign a value to
  // favoriteSnack
  
  // Print out the message
  // "My favorite snack is "
  // followed by the value in
  // favoriteSnack
  
}
```
- Assign a value to `favoriteSnack`, make it whatever snack is your favorite!
```
package main

import "fmt"

func main() {
  // Define a string variable
  // called favoriteSnack

  var favoriteSnack string

  // Assign a value to
  // favoriteSnack
  favoriteSnack = "crisps"
  
  // Print out the message
  // "My favorite snack is "
  // followed by the value in
  // favoriteSnack
  
}
```
- Using string concatenation, print out the message `"My favorite snack is "` followed by whatever the value inside of `favoriteSnack` is.
Be sure to add the trailing space to `"My favorite snack is "` or else it will render as `"My favorite snack isbroccoli"` instead of `"My favorite snack is broccoli"`.

```
package main

import "fmt"

func main() {
  // Define a string variable
  // called favoriteSnack

  var favoriteSnack string

  // Assign a value to
  // favoriteSnack
  favoriteSnack = "crisps"

  // Print out the message
  // "My favorite snack is "
  // followed by the value in
  // favoriteSnack
  fmt.Println("My favorite snack is " + favoriteSnack)
  
}
```
`Output`

```
My favorite snack is crisps
```


# Zero Values
Even before we assign anything to our variables they hold a value. Go’s designers attempted to make these “sensible defaults” that we can anticipate based on the variable’s types. All numeric variables have a value of `0` before assignment. String variables have a default value of `""`, which is also known as the empty string because it contains no characters. Boolean variables have a default value of `false`. For example:

```
var classTime uint32
var averageGrade float32
var teacherName string
var isPassFail bool
 
fmt.Println(classTime) // Prints 0
fmt.Println(averageGrade) // Prints 0
fmt.Println(teacherName) // Doesn't print anything
fmt.Println(isPassFail) // Prints false
```

Above we declared four variables, an unsigned 32-bit int, a 32-bit floating point number, a string, and a boolean. Without assigning any of the variables to a value we print them out to see their default value. The two numbers print the same result, `0`, a valid value for both types. The empty string, when printed, displays `nothing`. The boolean value prints `false`.


1. Create a variable emptyInt that’s of type int8.

`main.go`

```
package main

import "fmt"

func main() {
  // Create three variables
  // emptyInt an int8

  var emptyInt int8


  // emptyFloat a float32
  
  // and emptyString a string
  
  // Finally, print them all out
  
}
```

2. Create a variable called emptyFloat that’s of type float32.

`main.go`

```
package main

import "fmt"

func main() {
  // Create three variables
  // emptyInt an int8

  var emptyInt int8

  // emptyFloat a float32

  var emptyFloat float32
  
  // and emptyString a string
  
  // Finally, print them all out
  
}
```

3. Create an empty string called emptyString.

`main.go`

```
package main

import "fmt"

func main() {
  // Create three variables
  // emptyInt an int8
  var emptyInt int8

  // emptyFloat a float32  
  var emptyFloat float32
  
  // and emptyString a string
  var emptyString string
  
  // Finally, print them all out
  
}
```


4. Now print out all three variables on the same line. Since they’re not all strings you’ll have to separate them by commas, like so:

```
fmt.Println(emptyInt, emptyFloat, emptyString)
```

`main.go`

```
package main

import "fmt"

func main() {
  // Create three variables
  // emptyInt an int8
  var emptyInt int8
  // emptyFloat a float32
  var emptyFloat float32
  
  // and emptyString a string
  var emptyString string
  
  // Finally, print them all out
  fmt.Println(emptyInt, emptyFloat, emptyString)
  
}
```

`Output` after running `main.go`:

```
0 0 
```


# Inferring Type
There is a way to declare a variable without explicitly stating its type using the short declaration `:=` operator. We might use the `:=` operator if we know what value we want our variable to store when creating it. For instance:

```
nuclearMeltdownOccurring := true
radiumInGroundWater := 4.521
daysSinceLastWorkplaceCatastrophe := 0
externalMessage := "Everything is normal. Keep calm and carry on."
```

Above, we were able to define a `bool`, a `float`, an `int`, and a `string` without specifying the type. We used the `:=` to create a variable and infer its type based on the value provided. Floats created in this way are of type `float64`. Integers created in this way are either `int32` or `int64`.


Go also offers a separate syntax to declare a variable and infer its type. We could’ve written the same code from above as:

```
var nuclearMeltdownOccurring = true
var radiumInGroundWater = 4.521
var daysSinceLastWorkplaceCatastrophe = 0
var externalMessage = "Everything is normal. Keep calm and carry on."
```

Notice, in the second example, that we used the `var` keyword and the `=` operator. In both examples, we’ve declared and initialized variables while leaving the Go compiler to infer the type of the value assigned.


1. Create an integer called daysOnVacation using the := operator and assign it a number between 0 to 9.

`main.go`

```
package main

import "fmt"

func main() {
  // Define daysOnVacation using := below:
  daysOnVacation := 6
  
  // Define hoursInDay using var and = below:
  
  
  fmt.Println("You have spent", daysOnVacation * hoursInDay, "hours on vacation.")
}
```

2. Create another integer called hoursInDay with a value of 24, this time using var and =.

`main.go`

```
package main

import "fmt"

func main() {
  // Define daysOnVacation using := below:
  daysOnVacation := 6
  
  // Define hoursInDay using var and = below:
  var hoursInDay = 24
  
  fmt.Println("You have spent", daysOnVacation * hoursInDay, "hours on vacation.")
}
```

`Output` after running `main.go`:

```
You have spent 144 hours on vacation.
```


# Default int Type
There is one more common way to define an `int` in Go. Computers actually have a default length for the data in their Read-Only Memory (ROM). Some newer comps may have more processing power and can store/handle bigger chunks of data. These computers might be using a 64-bit architecture, but other computers still run on 32-bit architecture and work just fine. By providing the type `int` or `uint`, Go will check to see if the computer architecture is running on is 32-bit or 64-bit. Then it will either provide a 32-bit `int` (or `uint`) or a 64-bit one depending on the computer itself.

It’s recommended to use `int` unless there’s a reason to specify the size of the int (like knowing that value will be larger than the default type, or needing to optimize the amount of space used).


```
var timesWeWereFooled int
var foolishGamesPlayed uint
```

Above, we declared two variables, `timesWeWereFooled `an integer of either 32 or 64 bits. `foolishGamesPlayed`, an unsigned integer of either 32 or 64 bits.

```
consolationPrizes := 2
```

When a variable is declared and assigned a value using the `:=` operator, it will be the same type as if it were declared as an `int`. In the example above, `consolationPrize` has the type `int`.


1. Create an integer called cupsOfCoffeeConsumed. Make it an int type.

`main.go`

```
package main

import "fmt"

func main() {
  // Define cupsOfCoffeeConsumed here
  var cupsOfCoffeeConsumed int

  // Give a value to cupsOfCoffeeConsumed
  
  // Print out cupsOfCoffeeConsumed
  
}
```

2. Assign a value to cupsOfCoffeeConsumed. Make sure it’s a valid integer value!

`main.go`

```
package main

import "fmt"

func main() {
  // Define cupsOfCoffeeConsumed here
  var cupsOfCoffeeConsumed int
  // Give a value to cupsOfCoffeeConsumed
  var cupsOfCoffeeConsumed = 6
  // Print out cupsOfCoffeeConsumed
  
}
```

3. Print out the value of cupsOfCoffeeConsumed using fmt.Println()!

`main.go`

```
package main

import "fmt"

func main() {
  // Define cupsOfCoffeeConsumed here
  // var cupsOfCoffeeConsumed int

  // Give a value to cupsOfCoffeeConsumed
  var cupsOfCoffeeConsumed = 6

  // Print out cupsOfCoffeeConsumed
  fmt.Println(cupsOfCoffeeConsumed)
}
```

`Output` after running `main.go`:

```
6
```



