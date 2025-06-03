package service_test

import (
	"fmt"
	mock_user "go-hex-temp/internal/adapters/out/mock/user"
	"go-hex-temp/internal/core/domain"
	"go-hex-temp/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	mockRepo := new(mock_user.UserRepositoryMock)
	// qCom := service.NewQCompiler()
	userService := service.NewUserService(mockRepo, nil)

	t.Run("success", func(t *testing.T) {
		expectedUser := &domain.User{ID: "123", Name: "Test"}
		mockRepo.On("FindById", "123").Return(expectedUser, nil).Once()

		user, err := userService.GetUserById("123")
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("fail - user not found", func(t *testing.T) {
		mockRepo.On("FindById", "notfound").Return(nil, assert.AnError).Once()

		user, err := userService.GetUserById("notfound")
		assert.Error(t, err)
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})
}
func TestGetUsers(t *testing.T) {
	mockRepo := new(mock_user.UserRepositoryMock)
	qCom := service.NewQCompiler()
	userService := service.NewUserService(mockRepo, qCom)

	t.Run("success", func(t *testing.T) {
		query := domain.NewQuery()
		query.Filter["name"] = domain.Condition{
			domain.Eq: []any{"John"},
		}
		expectedUsers := []domain.User{
			{ID: "123", Name: "John"},
			{ID: "456", Name: "Developer"},
		}

		mockRepo.On("Find", query).Return(expectedUsers, nil).Once()

		user, err := userService.GetUsers(query)
		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("fail - error found", func(t *testing.T) {
		query := domain.NewQuery()
		mockRepo.On("Find", query).Return(nil, assert.AnError).Once()

		user, err := userService.GetUsers(query)
		assert.Error(t, err)
		fmt.Println(err, user)
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

}
