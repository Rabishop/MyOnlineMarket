package game

// UserRegistResponse struct
type gameUploadResponse struct {
	Status string `json:"status"`
}

// UserRegistRequest struct
type gameUploadRequest struct {
	GamePrice string `json:"gamePrice"`
	GameName  string `json:"gameName"`
	GameInfo  string `json:"userName"`
}
