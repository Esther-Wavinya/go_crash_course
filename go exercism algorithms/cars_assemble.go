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




// Task 2
// Calculate the number of working cars produced per minute

// Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute:

// rate := CalculateWorkingCarsPerMinute(1105, 90)
// // Output: 16
// Note: the return value should be an int.


// Calculate the number of working cars produced per minute
// Start by calculating the production of successful cars per hour. For this, you can use the CalculateProductionRatePerHour function you made from the previous step.

// Knowing the production per hour of cars, you can get the production per minute by dividing the production per hour by 60 (the number of minutes in an hour).

// Remember to cast the result to an int.


// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate ) / 60)   
}



// Task 3
// Calculate the cost of production

// Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. But with a bit of planning, 10 cars can be produced together for $95,000.

// For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars

// So the cost for 37 cars is: 3*95,000+7*10,000=355,000

// Implement the function CalculateCost that calculates the cost of producing a number of cars, regardless of whether they are successful:

// cost := CalculateCost(37)
// // Output: 355000

// cost = CalculateCost(21)
// // Output: 200000
// Note: the return value should be an uint.


// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
    return((uint(carsCount / 10) * 95000) + (uint(carsCount % 10) * 10000 ))    
}