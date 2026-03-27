package v1

import "github.com/Oleska1601/WBOptimizeServer/internal/models"

type ServiceI interface {
	Fibonacci(n int) int
	ProcessJSON(item models.ItemV1) ([]byte, error)
}
