package auth

type LoginDTO struct {
	Username string
	Password string
}

type SignupDTO struct {
	Username string
	Email    string
	Password string
}
