package api

import (
	"net/http"
	"rag-app/internal/services"

	"github.com/gin-gonic/gin"
)


// Add this handler definition
func IngestHandler(c *gin.Context) {
	var doc DocumentRequest
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.ProcessDocument(doc.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Document ingested successfully"})
}

func QueryHandler(c *gin.Context) {
	var query QueryRequest
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve relevant context for the query
	context, err := services.RetrieveContext(query.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve context"})
		return
	}

	// Generate an answer using the retrieved context and question
	answer, err := services.GenerateAnswer(query.Question, context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate answer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": answer})
}
