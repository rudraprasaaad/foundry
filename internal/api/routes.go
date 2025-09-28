package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/login", LoginHandler)
		v1.GET("/metrics/history", MetricsHistoryHandler)
		v1.GET("/metrics/live", LiveMetricsHandler)
	}
}
