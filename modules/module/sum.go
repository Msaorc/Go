package module

func Sum[T int | float32](a, b T) T {
	return a + b
}
