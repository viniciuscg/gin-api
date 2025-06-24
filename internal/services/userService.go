package services

import (
	"github.com/viniciuscg/gin-api/internal/model"
	"github.com/viniciuscg/gin-api/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	users, err := s.Repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
