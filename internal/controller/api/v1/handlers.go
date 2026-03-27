package v1

import (
	"math/rand/v2"
	"net/http"
	"net/http/pprof"
	"strconv"
	"time"

	"github.com/Oleska1601/WBOptimizeServer/internal/models"
	"github.com/gin-gonic/gin"
)

const (
	cpuBoundURI    = "/cpu"
	memoryBoundURI = "/memory"
)

const (
	pprofGroupURI = "/debug/pprof"

	profileURI = "/profile"
	traceURI   = "/trace"
	heapURI    = "/heap"
	allocsURI  = "/allocs"
)

func (v1 *APIV1) registerHandlers(group *gin.RouterGroup) {
	group.GET(cpuBoundURI, v1.cpuBound)
	group.GET(memoryBoundURI, v1.memoryBound)

	pprofGroup := group.Group(pprofGroupURI)
	pprofGroup.GET(profileURI, gin.WrapF(pprof.Profile))
	pprofGroup.GET(traceURI, gin.WrapF(pprof.Trace))
	pprofGroup.GET(heapURI, gin.WrapH(pprof.Handler("heap")))
	pprofGroup.GET(allocsURI, gin.WrapH(pprof.Handler("allocs")))
}

func (v1 *APIV1) cpuBound(c *gin.Context) {
	n := rand.IntN(20) + 1
	start := time.Now()
	result := v1.service.Fibonacci(n)
	duration := time.Since(start)

	v1.logger.Info().
		Str("path", "apiV1 cpuBound").
		Int("n", n).
		Int("result", result).
		Dur("duration", duration).
		Msg("fibonacci v1 calculated")

	c.JSON(http.StatusOK, gin.H{
		"n":           n,
		"result":      result,
		"duration_ms": duration.Milliseconds(),
	})
}

func (v1 *APIV1) memoryBound(c *gin.Context) {
	req := models.ItemV1{
		Active:    true,
		SKU:       strconv.Itoa(rand.Int()),
		Price:     rand.Float64(),
		Qty:       rand.Int32(),
		Name:      strconv.Itoa(rand.Int()),
		Weight:    float32(rand.Float64()),
		Width:     rand.Int32(),
		Height:    rand.Int32(),
		Depth:     rand.Int32(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	start := time.Now()
	result, err := v1.service.ProcessJSON(req)
	if err != nil {
		v1.logger.Error().Err(err).Msg("failed to process JSON")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	duration := time.Since(start)
	v1.logger.Info().
		Str("sku", req.SKU).
		Dur("duration", duration).
		Int("result_size", len(result)).
		Msg("JSON processed")

	c.Data(http.StatusOK, "application/json", result)
}
