package api

type DocumentRequest struct {
	Content string `json:"content" binding:"required"`
}

type QueryRequest struct {
	Question string `json:"question" binding:"required"`
}
