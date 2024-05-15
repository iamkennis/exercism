package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100

	// panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	hrs := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(hrs) / 60
	// panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	eachCarCost := 10000
	tenCarCost := 95000

	groupsOfTen := carsCount / 10

	remainingCars := carsCount % 10

	totalCost := groupsOfTen*tenCarCost + remainingCars*eachCarCost

	return uint(totalCost)
	// panic("CalculateCost not implemented")
}
