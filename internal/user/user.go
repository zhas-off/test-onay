package user

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Comment - структура для наших пользователей
type User struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

// Store - интерфейс, который нам нужен для хранения пользователей
type Store interface {
	GetUsers(*fiber.Ctx) (User, error)
	PostUser(*fiber.Ctx, User) (User, error)
	UpdateUser(*fiber.Ctx, string, User) (User, error)
}

// Service - структура для нашего сервиса пользователей
type Service struct {
	Store Store
}

// NewService - возвращает новый сервис для пользователей
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetUsers(ctx *fiber.Ctx) (User, error) {
	users, err := s.Store.GetUsers(ctx)
	if err != nil {
		log.Errorf("an error occured fetching the comment: %s", err.Error())
		return User{}, errors.New("could not fetch comment by ID")
	}
	return users, nil
}

func (s *Service) PostUser(ctx *fiber.Ctx, user User) (User, error) {
	user, err := s.Store.PostUser(ctx, user)
	if err != nil {
		log.Errorf("an error occurred adding the comment: %s", err.Error())
	}
	return user, nil
}

func (s *Service) UpdateUser(
	ctx *fiber.Ctx, ID string, newUser User,
) (User, error) {
	user, err := s.Store.UpdateUser(ctx, ID, newUser)
	if err != nil {
		log.Errorf("an error occurred updating the comment: %s", err.Error())
	}
	return user, nil
}
