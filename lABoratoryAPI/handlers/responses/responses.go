package responses

type Response struct {
	Message string `json:"response"`
}

type ResponseWithToken struct {
	Message string `json:"response"`
	Token   string `json:"token"`
}

type ResponseWithError struct {
	Message string `json:"response"`
	Error   string `json:"error"`
}

type DeleteResponse struct {
	WasDeleted bool `json:"wasDeleted"`
}
