package utils

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func ArgMax(inputs []float64) (float64, int) {
	max := inputs[0]
	maxIndex := 0
	for i, num := range inputs {
		if num > max {
			max = num
			maxIndex = i
		}
	}
	return max, maxIndex
}

func ArgMin(inputs []float64) (float64, int) {
	min := inputs[0]
	minIndex := 0
	for i, num := range inputs {
		if num < min {
			min = num
			minIndex = i
		}
	}
	return min, minIndex
}

// random float between -1 and 1
func Random() float64 {
	return 2*rand.Float64() - 1
}

func Lerp(x, y, t float64) float64 {
	return x + (y-x)*t
}

func AddMaps(m1, m2 map[string]float64) map[string]float64 {
	result := make(map[string]float64)
	for k, v := range m1 {
		result[k] += v
	}
	for k, v := range m2 {
		result[k] += v
	}
	return result
}

func BinarySearch(arr []string, target string) int {

	//slices.Sort(arr)

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func NewRandomDense(rows, cols int) *mat.Dense {
	data := make([]float64, rows*cols)
	for i := range data {
		data[i] = Random()
	}
	return mat.NewDense(rows, cols, data)
}

func PrintMatrix(m *mat.Dense) {
	rows, cols := m.Dims()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%.4f\t", m.At(i, j))
		}
		fmt.Println()
	}
}
