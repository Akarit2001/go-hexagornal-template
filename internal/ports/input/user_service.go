package input

import "go-hex-temp/internal/core/domain"

type UserService interface {
	Create(user *domain.User) error
	GetUserById(userId string) (*domain.User, error)
	GetUsers(q *domain.Query) ([]domain.User, error)
}
