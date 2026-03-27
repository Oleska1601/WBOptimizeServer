package v1service

import (
	"bytes"
	"encoding/json"

	"github.com/Oleska1601/WBOptimizeServer/internal/models"
)

type ServiceV1 struct {
}

func New() *ServiceV1 {
	return &ServiceV1{}
}

// cpu bound
func (s *ServiceV1) Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return s.Fibonacci(n-1) + s.Fibonacci(n-2)
}

// memory bound
func (s *ServiceV1) ProcessJSON(item models.ItemV1) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(item); err != nil {
		return nil, err
	}

	result := make([]byte, buffer.Len())
	copy(result, buffer.Bytes())
	return result, nil
}
