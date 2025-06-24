package service

import (
	"go-hex-temp/internal/core/domain"
	"go-hex-temp/internal/infrastructure/logx"
	"go-hex-temp/internal/ports/input"
	"go-hex-temp/internal/ports/output"
)

type userService struct {
	userRepo  output.UserRepository
	cache     typedCache[domain.User]
	qCompiler *QCompiler
}

func NewUserService(userRepo output.UserRepository, qCompiler *QCompiler, cache output.Cache) input.UserService {

	return &userService{
		userRepo:  userRepo,
		qCompiler: qCompiler,
		cache:     typedCache[domain.User]{cache},
	}
}

// Create implements input.UserService.
func (u *userService) Create(user *domain.User) error {
	_, err := u.userRepo.Save(user)
	if err != nil {

		logx.Error(err.Error())
		return err
	}
	return nil
}

func (u *userService) GetUserById(userId string) (*domain.User, error) {
	cb := func() (*domain.User, error) {
		user, err := u.userRepo.FindById(userId)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	key := "user:" + userId

	user, err := u.cache.ReadOne(key, 30, cb)
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userService) GetUsers(q *domain.Query) (*domain.Paginated[domain.User], error) {

	compiledQuery, err := u.qCompiler.Compile(q, domain.User{})
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}

	return u.userRepo.Find(compiledQuery)
}
