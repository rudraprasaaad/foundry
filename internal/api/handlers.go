package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "",
	})
}

func MetricsHistoryHandler(c *gin.Context) {
	mockHistory := []gin.H{
		{"serviceName": "api-1", "cpuUsage": 65.5, "memoryUsage": 280, "timestamp": "2025-09-28T18:00:00Z"},
		{"serviceName": "worker-1", "cpuUsage": 72.1, "memoryUsage": 310, "timestamp": "2025-09-28T18:01:00Z"},
		{"serviceName": "api-1", "cpuUsage": 68.3, "memoryUsage": 295, "timestamp": "2025-09-28T18:02:00Z"},
	}
	c.JSON(http.StatusOK, mockHistory)
}

func LiveMetricsHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.Request.Context().Done():
			return
		case t := <-ticker.C:
			data := fmt.Sprintf(`data: {"cpuUsage": %.2f, "memoryUsage": %.2f, "timestamp": "%s"}\n\n`, 60+10*time.Since(t).Seconds(), 300+20*time.Since(t).Seconds(), t.Format(time.RFC3339))
			c.Writer.Write([]byte(data))
			c.Writer.Flush()
		}
	}
}
