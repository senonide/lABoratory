package responses

type Response struct {
	Message string `json:"response"`
}

type ResponseWithError struct {
	Message string `json:"response"`
	Error   string `json:"error"`
}

type DeleteResponse struct {
	WasDeleted bool `json:"wasDeleted"`
}
