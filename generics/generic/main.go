package main

// Sum[T int | float64] Generic

func Sum[T int | float64](m map[int]T) T {
	var result T

	for _, v := range m {
		result += v
	}

	return result
}

func main() {
	x := map[int]int{1: 500, 2: 1500, 3: 2500}
	y := map[int]float64{1: 500.00, 2: 1500.00, 3: 2500.00}

	println(Sum(x))
	println(Sum(y))
}
