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

// Output after running code

// Current station: Union Square
// Next train: 12 minutes
// Is uptown: false

// Current station: Grand Central
// Next train: 3 minutes
// Is uptown: true
