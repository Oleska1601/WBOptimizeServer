package v2

import "github.com/Oleska1601/WBOptimizeServer/internal/models"

type ServiceI interface {
	Fibonacci(n int) int
	ProcessJSON(item *models.ItemV2) ([]byte, error)
}
