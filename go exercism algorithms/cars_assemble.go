// Instructions
// In this exercise you'll be writing code to analyze the production in a car factory.

// Task 1
// Calculate the number of working cars produced per hour

// The cars are produced on an assembly line. The assembly line has a certain speed, that can be changed. The faster the assembly line speed is, the more cars are produced. However, changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.

// Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as a percentage, from 0 to 100:

// rate := CalculateWorkingCarsPerHour(1547, 90)
// // Output: 1392.3
// Note: the return value should be a float64.

// Calculate the number of working cars produced per hour
// The percentage (passed as an argument) is a number between 0-100. To make this percentage a bit easier to work with, start by dividing it by 100.

// To compute the number of cars produced successfully, multiply the percentage (divided by 100) by the number of cars produced.

// When multiplying two numbers together, they both need to be of the same type. Use type conversions if needed.

package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return((successRate / 100 ) * float64(productionRate)) 
}