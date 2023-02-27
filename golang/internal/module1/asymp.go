package module1

import (
	"math"
)

func Alg1time(n, m int) int {
	return int(math.Log(float64(n))*math.Log(float64(m))) + n*n*m + 5*n
}
func Alg2time(n, m int) int {
	return 40*n + 30*m + 1500
}
