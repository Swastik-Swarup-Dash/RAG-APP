package api

import (
    "net/http"
	"rag-app/internal/services"

	"github.com/gin-gonic/gin"
    
 )


fun IngestHandler(c *gin.Context) {
    var doc DocumentRequest
    if err := c.ShouldBindJSON(&doc); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.ProcessDocument(doc.Content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process document"})    
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Document ingested successfully"})
}


func QuerHandler(c, *gin.Context) {
    var query QueryRequest
    if err := c.ShouldBindJSON(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context, err := services.RetrieveContext(query.Question)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query document"})
        return
    }


    answer, err := services.GenerateAnswer(query.Question, context)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate answer"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"answer": answer})
}
