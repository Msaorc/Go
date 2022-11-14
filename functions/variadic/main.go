package main

func main() {
	println(sum(1, 35, 84, 95, 86, 74, 25, 12, 36, 35, 25, 42, 49, 57, 10, 87, 522, 3684, 21, 5777, 569, 36, 21, 47))
	println(sum(30, 28, 73, 68))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
