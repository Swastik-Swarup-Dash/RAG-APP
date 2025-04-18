package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/ingest", IngestHandler)
	r.POST("/query", QueryHandler)

	return r
}