package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
    r.GET("/health", HealthHandler)
	r.POST("/ingest", IngestHandler)
	r.POST("/query", QueryHandler)
    r.Run(":8080")
	return r
}
