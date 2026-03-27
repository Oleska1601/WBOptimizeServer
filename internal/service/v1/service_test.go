package v1service

import (
	"testing"

	"github.com/Oleska1601/WBOptimizeServer/internal/models"
)

// CPU-bound бенчмарки
func BenchmarkFibonacci(b *testing.B) {
	s := New()
	n := 30

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Fibonacci(n)
	}
}

// Memory-bound бенчмарки
func BenchmarkProcessJSON(b *testing.B) {
	s := New()

	// Подготовка тестовых данных
	item := models.ItemV1{
		SKU:    "PROD-001",
		Price:  99.99,
		Qty:    5,
		Name:   "Test Product",
		Weight: 1.5,
		Width:  10,
		Height: 20,
		Depth:  30,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := s.ProcessJSON(item)
		if err != nil {
			b.Fatal(err)
		}
	}
}
