package postgresql

type Repositories struct {
	User *UserRepository
}

func NewRepositories() *Repositories {
	return &Repositories{User: NewUserRepository()}
}
