package main

// 5º Ejercicio
// Agregar un test, un benchmark y un example al ejercicio de suma de un valor parametrizado p
// entre 0 <= v <= p.

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	slice, total := calculate(10)
	if len(slice) == 0 {
		t.Error("Expected 4, got ", slice)
	}
	if total == 0 {
		t.Error("Expected 23, got ", total)
	}
}

func TestCalculateTwo(t *testing.T) {
	slice, total := calculate(5)
	if slice[0] != 3 {
		t.Error("Expected 3, got ", slice)
	}
	if total != 3 {
		t.Error("Expected 3, got ", total)
	}
}

func BenchmarkCalculate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculate(10)
	}
}

func ExampleCalculateOne() {
	slice, _ := calculate(5)
	fmt.Println(slice)
	// Output: [3]
}

func ExampleCalculateTwo() {
	_, total := calculate(10)
	fmt.Println(total)
	// Output: 23
}
