package v1

import (
	"github.com/Oleska1601/WBOptimizeServer/pkg/logger"
	"github.com/gin-gonic/gin"
)

const (
	V1GroupURI = "/v1"
)

type APIV1 struct {
	service ServiceI
	logger  logger.LoggerI
}

func New(service ServiceI, logger logger.LoggerI) *APIV1 {
	return &APIV1{
		service: service,
		logger:  logger,
	}
}

func (v1 *APIV1) RegisterHandlers(group *gin.RouterGroup) {
	v1Group := group.Group(V1GroupURI)

	v1.registerHandlers(v1Group)
}
