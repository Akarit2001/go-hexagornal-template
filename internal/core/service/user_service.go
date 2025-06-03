package service

import (
	"go-hex-temp/internal/core/domain"
	"go-hex-temp/internal/infrastructure/logx"
	"go-hex-temp/internal/ports/input"
	"go-hex-temp/internal/ports/output"
)

type userService struct {
	userRepo  output.UserRepository
	qCompiler *QCompiler
}

func NewUserService(userRepo output.UserRepository, qCompiler *QCompiler) input.UserService {

	return &userService{
		userRepo:  userRepo,
		qCompiler: qCompiler,
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

	user, err := u.userRepo.FindById(userId)
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (u *userService) GetUsers(q *domain.Query) ([]domain.User, error) {

	compiledQuery, err := u.qCompiler.Compile(q, domain.User{})
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}

	return u.userRepo.Find(compiledQuery)
}
