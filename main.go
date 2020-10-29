package main

import "fmt"

func main() {
	n := []float32{120.4, -46.7, 32.50, 34.65, -67.45}
	fmt.Printf("The total is %.02f\n", sum(n))
}

func sum(s []float32) float32 {
	var t float32
	for _, v := range s {
		if t < 0 {
			t = add(t, v)
		} else {
			t = sub(t, v)
		}
	}

	return t
}

//go:inline
func add(a, b float32) float32 {
	if b < 0 {
		panic(`Do not add negative number`)
	}
	return a + b
}

//go:inline
func sub(a, b float32) float32 {
	return a - b
}
