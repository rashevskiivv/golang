package mymath

func Avg(values []float64) float64 {
	total := 0.0
	for _, x := range values {
		total += x
	}
	return total / float64(len(values))
}
