package service_test

import (
	"fmt"
	"go-hex-temp/internal/adapters/out/repositories"
	"go-hex-temp/internal/core/domain"
	"go-hex-temp/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	mockRepo := new(repositories.UserRepositoryMock)
	// qCom := service.NewQCompiler()
	userService := service.NewUserService(mockRepo, nil, nil)

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
	mockRepo := new(repositories.UserRepositoryMock)
	qCom := service.NewQCompiler()
	userService := service.NewUserService(mockRepo, qCom, nil)

	t.Run("success", func(t *testing.T) {
		query := domain.NewQuery()
		query.Filter["name"] = domain.QCondition{
			domain.Eq: []any{"John"},
		}
		expectedUsers := []domain.User{
			{ID: "123", Name: "John"},
			{ID: "456", Name: "Developer"},
		}

		expectedPage := &domain.Paginated[domain.User]{
			Items:      expectedUsers,
			TotalCount: 30,
		}

		mockRepo.On("Find", query).Return(expectedPage, nil).Once()

		page, err := userService.GetUsers(query)
		assert.NoError(t, err)
		assert.Equal(t, expectedPage, page)

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
