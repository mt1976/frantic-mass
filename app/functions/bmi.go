package functions

func CalculateBMI(heightCm float64, weightKg float64) float64 {
	heightM := heightCm / 100.0
	thisBMI := weightKg / (heightM * heightM)
	return thisBMI
}
