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

const OvenTime = 40


// 2. Calculate the remaining oven time in minutes
// Define the RemainingOvenTime() function that takes the actual minutes the lasagna has been in the oven as a parameter and returns how many minutes the lasagna still has to remain in the oven, based on the expected oven time in minutes from the previous task.

// func RemainingOvenTime(actual int) int {
//     // TODO
// }

// RemainingOvenTime(30)
// // Output: 10
// 3. Calculate the preparation time in minutes
// Define the PreparationTime function that takes the number of layers you added to the lasagna as a parameter and returns how many minutes you spent preparing the lasagna, assuming each layer takes you 2 minutes to prepare.

// func PreparationTime(numberOfLayers int) int {
//     // TODO
// }

// PreparationTime(2)
// // Output: 4
// 4. Calculate the elapsed working time in minutes
// Define the ElapsedTime function that takes two parameters: the first parameter is the number of layers you added to the lasagna, and the second parameter is the number of minutes the lasagna has been in the oven. The function should return how many minutes in total you've worked on cooking the lasagna, which is the sum of the preparation time in minutes, and the time in minutes the lasagna has spent in the oven at the moment.

// func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
//     // TODO
// }

// ElapsedTime(3, 20)
// // Output: 26