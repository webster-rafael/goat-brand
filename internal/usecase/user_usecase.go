package usecase

import "goatbrand-backend/internal/domain"

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUsers() ([]domain.User, error) {
	return u.repo.FindAll()
}
