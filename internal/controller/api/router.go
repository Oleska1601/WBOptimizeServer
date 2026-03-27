package api

import (
	"github.com/Oleska1601/WBOptimizeServer/config"
	"github.com/gin-gonic/gin"
)

const (
	APIGroupURI = "/api"
)

type HTTPController interface {
	RegisterHandlers(*gin.RouterGroup)
}

func Register(cfg *config.GinConfig, controller HTTPController) *gin.Engine {
	engine := gin.New()
	gin.SetMode(cfg.Mode)

	group := engine.Group(APIGroupURI)
	controller.RegisterHandlers(group)
	return engine
}
