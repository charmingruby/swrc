package dto

type RegisterInputDTO struct {
	GithubDisplayName string
	Email             string
	Password          string
}

type RegisterOutputDTO struct {
	ID       string
	Role     string
	IsValid  bool
	Verified bool
}
