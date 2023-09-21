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

// PostCommentRequest
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

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	comment, err := h.Service.GetUsers(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(comment)
}

func (h *Handler) PostUser(c *fiber.Ctx) error {
	var postUserReq PostUserRequest

	if err := c.BodyParser(&postUserReq); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	cmt := userFromPostUserRequest(postUserReq)
	cmt, err := h.Service.PostUser(c, cmt)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(cmt)
}

// UpdateComment - обновить комментарий
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	commentID := c.Params("id")

	var updateUserRequest UpdateUserRequest
	if err := c.BodyParser(&updateUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	cmt := userFromUpdateUserRequest(updateUserRequest)

	cmt, err := h.Service.UpdateUser(c, commentID, cmt)
	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(cmt)
}
