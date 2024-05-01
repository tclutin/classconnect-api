package http

type APIErrorResponse struct {
	Error any `json:"message"`
}

func NewAPIResponse(message any) APIErrorResponse {
	return APIErrorResponse{Error: message}
}
