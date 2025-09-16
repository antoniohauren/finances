package models

type ErrorResponse struct {
	Reason string `json:"reason"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
