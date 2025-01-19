package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/toastsandwich/networking-tic-tac-toe/model"
	"github.com/toastsandwich/networking-tic-tac-toe/repository"
)

type Service struct {
	Repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) CreateUserService(user *model.User) error {
	email := []byte(user.Email)

	buf, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return s.Repository.InsertUser(email, buf)
}

func (s *Service) GetUserService(email string) (*model.User, error) {
	data, err := s.Repository.GetUser([]byte(email))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("user not found")
	}
	user := &model.User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) DeleteUserService(email string) error {
	return s.Repository.DeleteUser([]byte(email))
}

func (s *Service) AddTokenToBlackListService(token string) error {
	timenow := time.Now().Format(time.RFC3339)
	return s.Repository.AddTokenToBlackList([]byte(token), []byte(timenow))
}

func (s *Service) CheckForBlacklistService(token string) error {
	return s.Repository.FindToken([]byte(token))
}
