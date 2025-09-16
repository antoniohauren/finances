package models

type AuthSignUpRequest CreateUserRequest
type AuthSignUpResponse CreateUserResponse

type AuthSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSignInResponse CreateUserResponse
