package mock_user

import (
	"go-hex-temp/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

// Find mocks searching for users with a query.
func (u *UserRepositoryMock) Find(query *domain.Query) ([]domain.User, error) {
	args := u.Called(query)
	if users, ok := args.Get(0).([]domain.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

// FindById mocks finding a user by ID.
func (u *UserRepositoryMock) FindById(id string) (*domain.User, error) {
	args := u.Called(id)
	if user, ok := args.Get(0).(*domain.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

// Save mocks saving a user and returning the result.
func (u *UserRepositoryMock) Save(user *domain.User) (*domain.User, error) {
	args := u.Called(user)
	if savedUser, ok := args.Get(0).(*domain.User); ok {
		return savedUser, args.Error(1)
	}
	return nil, args.Error(1)
}
