package v2service

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/Oleska1601/WBOptimizeServer/internal/models"
)

type ServiceV2 struct {
	pool *sync.Pool
}

func New() *ServiceV2 {
	return &ServiceV2{
		pool: &sync.Pool{
			New: func() any {
				return bytes.NewBuffer(make([]byte, 0, 1024))
			},
		},
	}
}

// cpu bound
func (s *ServiceV2) Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// memory bound
func (s *ServiceV2) ProcessJSON(item *models.ItemV2) ([]byte, error) {
	buffer := s.pool.Get().(*bytes.Buffer)
	defer s.pool.Put(buffer)

	buffer.Reset()

	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(item); err != nil {
		return nil, err
	}

	result := make([]byte, buffer.Len())
	copy(result, buffer.Bytes())
	return result, nil
}
