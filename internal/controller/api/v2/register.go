package v2

import (
	"github.com/Oleska1601/WBOptimizeServer/pkg/logger"
	"github.com/gin-gonic/gin"
)

const (
	V2GroupURI = "/v2"
)

type APIV2 struct {
	service ServiceI
	logger  logger.LoggerI
}

func New(service ServiceI, logger logger.LoggerI) *APIV2 {
	return &APIV2{
		service: service,
		logger:  logger,
	}
}

func (v2 *APIV2) RegisterHandlers(group *gin.RouterGroup) {
	v2Group := group.Group(V2GroupURI)

	v2.registerHandlers(v2Group)
}
