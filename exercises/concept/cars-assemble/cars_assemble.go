package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {

	switch {
	case speed == 0:
		return 0

	case speed <= 4 && speed >= 1:
		return 1
	case 5 <= speed && speed <= 8:
		return 0.9
	case speed == 9 || speed == 10:
		return 0.77
	default:
		panic("not supported, 0 ~ 10 only")
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	if speed < 0 || speed > 10 {
		panic("not supported, 0 ~ 10 only")
	}

	return SuccessRate(speed) * 221 * float64(speed)

}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	total := CalculateProductionRatePerHour(speed)
	return int(total) / 60
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	total := CalculateProductionRatePerHour(speed)
	if total <= limit {
		return total
	} else {
		return limit
	}

}
