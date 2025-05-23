package system

import "github.com/gin-gonic/gin"

type SystemHandler struct{}

// HealthCheck godoc
// @Summary 死活監視用
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /v1/health [get]
func (handler *SystemHandler) Health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Status": "ok",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
