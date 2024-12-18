package statistics

// Avg возвращает среднее значение массива чисел
func Avg(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}
