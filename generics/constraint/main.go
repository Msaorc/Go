package main

type Number int

type ContraintNumber interface {
	// ~ to accept primitive type
	~int | ~float64
}

func Sum[T ContraintNumber](m map[int]T) T {
	var result T

	for _, v := range m {
		result += v
	}

	return result
}

func main() {
	x := map[int]int{1: 500, 2: 1500, 3: 2500}
	y := map[int]float64{1: 500.00, 2: 1500.00, 3: 2500.00}
	z := map[int]Number{1: 500, 2: 1500, 3: 5000}

	println(Sum(x))
	println(Sum(y))
	println(Sum(z))
}
