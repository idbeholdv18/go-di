package service

import "github/idbeholdv18/go-di/internal/repository"

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(db repository.UserRepository) *UserService {
	return &UserService{
		repo: db,
	}
}

func (s *UserService) Create(name string, age int) *repository.User {
	return s.repo.Create(name, age)
}

func (s *UserService) Delete(name string) *repository.User {
	return s.repo.DeleteByName(name)
}

func (s *UserService) Find(name string) *repository.User {
	return s.repo.FindByName(name)
}
