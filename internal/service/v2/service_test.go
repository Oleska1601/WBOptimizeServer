package v2service

import (
	"testing"

	"github.com/Oleska1601/WBOptimizeServer/internal/models"
)

func BenchmarkFibonacci(b *testing.B) {
	s := New()
	n := 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Fibonacci(n)
	}
}

func BenchmarkProcessJSON(b *testing.B) {
	s := New()

	// Подготовка тестовых данных (указатель!)
	item := &models.ItemV2{
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
