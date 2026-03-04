package usecase

import (
	"errors"
	"goatbrand-backend/internal/domain"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUsers() ([]domain.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecase) Login(email, code string) (*domain.User, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Usuário não cadastrado.")
	}
	if user.Code != code {
		return nil, errors.New("Código incorreto.")
	}
	return user, nil
}
