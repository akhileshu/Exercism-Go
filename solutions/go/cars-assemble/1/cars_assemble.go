package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
//
// Type conversion:
// productionRate is int, but successRate is float64.
// Convert productionRate to float64 so both values can participate
// in floating-point multiplication.
//
// Example:
// productionRate = 221
// successRate    = 90.0
//
// 90.0 / 100 = 0.9
// float64(221) = 221.0
// 0.9 * 221.0 = 198.9
//
// Output: 198.9
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return successRate / 100 * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
//
// Type conversion:
// CalculateWorkingCarsPerHour returns float64.
// Dividing by 60 can still produce a fractional value.
// The result must be an int, so convert the final float64 to int.
//
// Converting float64 -> int discards the fractional part.
//
// Example:
// workingCarsPerHour = 198.9
//
// 198.9 / 60 = 3.315
// int(3.315) = 3
//
// Output: 3
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
//
// Type conversion:
// batches and singlePieces are ints because carsCount is an int.
// The arithmetic therefore produces an int.
//
// The function must return uint, so the final result is converted
// from int to uint.
//
// Example:
// carsCount = 13
//
// batches      = 13 / 10 = 1
// singlePieces = 13 % 10 = 3
//
// cost = 1*95000 + 3*10000
//      = 125000
//
// int(125000) -> uint(125000)
//
// Output: 125000
func CalculateCost(carsCount int) uint {
	batches := carsCount / 10
	singlePieces := carsCount % 10

	return uint(batches*95000 + singlePieces*10000)
}