package output

import "go-hex-temp/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindById(id string) (*domain.User, error)
	Find(query *domain.Query) ([]domain.User, error)
}
