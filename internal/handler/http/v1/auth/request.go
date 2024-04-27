package auth

import "classconnect-api/internal/domain/auth"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l LoginRequest) ToDTO() auth.LoginDTO {
	return auth.LoginDTO{
		Username: l.Username,
		Password: l.Password,
	}
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l SignupRequest) ToDTO() auth.SignupDTO {
	return auth.SignupDTO{
		Username: l.Username,
		Email:    l.Email,
		Password: l.Password,
	}
}
