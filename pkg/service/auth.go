package service

import (
	timetracker "time-tracker"
	"time-tracker/pkg/repository"
)

type AuthServise struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthServise {
	return &AuthServise{repo: repo}
}

func (s *AuthServise) CreateUser(user timetracker.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthServise) GetAllUsers() ([]timetracker.User, error) {
	return s.repo.GetAllUsers()
}

func (s *AuthServise) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *AuthServise) UpdateUser(id int, input timetracker.UpdateUserInput) error {
	return s.repo.UpdateUser(id, input)
}

func (s *AuthServise) Search(nameFilter, passportNumberFilter string, page, pageSize int) ([]timetracker.User, error) {
	return s.repo.Search(nameFilter, passportNumberFilter, page, pageSize)
}
