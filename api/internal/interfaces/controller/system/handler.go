package system

import "github.com/gin-gonic/gin"

type SystemHandler struct{}

type Response struct {
	Status  string `json:"status"`  // ステータス（成功、失敗など）
	Message string `json:"message"` // メッセージ（詳細説明）
}

// HealthCheck godoc
// @Summary 死活監視用
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} system.Response "成功時のレスポンス"
// @Router /v1/health [get]
func (h *SystemHandler) Health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Status": "ok",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
