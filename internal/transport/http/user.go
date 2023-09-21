package http

import (
	log "github.com/sirupsen/logrus"
	"github.com/zhas-off/test-onay/internal/user"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetUsers(ctx *fiber.Ctx) (user.User, error)
	PostUser(ctx *fiber.Ctx, usr user.User) (user.User, error)
	UpdateUser(ctx *fiber.Ctx, ID string, newUser user.User) (user.User, error)
}

type PostUserRequest struct {
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

func userFromPostUserRequest(u PostUserRequest) user.User {
	return user.User{
		FullName: u.FullName,
		Age:      u.Age,
		Address:  u.Address,
	}
}

type UpdateUserRequest struct {
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

func userFromUpdateUserRequest(u UpdateUserRequest) user.User {
	return user.User{
		FullName: u.FullName,
		Age:      u.Age,
		Address:  u.Address,
	}
}

// GetUser - получает и показывает всех пользователей
func (h *Handler) GetUsers(c *fiber.Ctx) error {
	user, err := h.Service.GetUsers(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(user)
}

// PostUser - добавляет нового пользователя
func (h *Handler) PostUser(c *fiber.Ctx) error {
	var postUserReq PostUserRequest

	if err := c.BodyParser(&postUserReq); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	user := userFromPostUserRequest(postUserReq)
	user, err := h.Service.PostUser(c, user)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(user)
}

// UpdateUser - обновляет данные пользователя
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var updateUserRequest UpdateUserRequest
	if err := c.BodyParser(&updateUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	user := userFromUpdateUserRequest(updateUserRequest)

	user, err := h.Service.UpdateUser(c, userID, user)
	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(user)
}
