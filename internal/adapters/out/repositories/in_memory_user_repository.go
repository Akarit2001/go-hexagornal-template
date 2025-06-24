package repositories

import (
	"errors"

	"go-hex-temp/internal/core/domain"
)

type inMemoryRepoUser struct {
	users map[string]*domain.User
}

func NewInMemoryRepoUser() *inMemoryRepoUser {
	return &inMemoryRepoUser{
		users: map[string]*domain.User{
			"1": {ID: "1", Name: "Alice", Email: "alice@example.com", Password: "hashed1", Age: 30},
			"2": {ID: "2", Name: "Bob", Email: "bob@example.com", Password: "hashed2", Age: 25},
			"3": {ID: "3", Name: "Charlie", Email: "charlie@example.com", Password: "hashed3", Age: 35},
		},
	}
}

func (r *inMemoryRepoUser) Save(user *domain.User) (*domain.User, error) {
	return nil, errors.New("read-only repository: Save not implemented")
}

func (r *inMemoryRepoUser) FindById(id string) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *inMemoryRepoUser) Find(query *domain.Query) (*domain.Paginated[domain.User], error) {
	// Just return all users, ignoring the query

	var allUsers []domain.User
	for _, u := range r.users {
		allUsers = append(allUsers, *u)
	}

	return &domain.Paginated[domain.User]{
		Items:      allUsers,
		TotalCount: int64(len(allUsers)),
	}, nil
}
