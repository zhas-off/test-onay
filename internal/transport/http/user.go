package http

import (
	log "github.com/sirupsen/logrus"
	"github.com/zhas-off/test-onay/internal/user"

	"github.com/gofiber/fiber/v2"
)

type CommentService interface {
	GetUsers(ctx *fiber.Ctx) (user.User, error)
	PostUser(ctx *fiber.Ctx, usr user.User) (user.User, error)
	UpdateUser(ctx *fiber.Ctx, ID string, newUser user.User) (user.User, error)
}

// // GetComment - get and show a comment by ID
// func (h *Handler) GetUsers(c *fiber.Ctx) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	usr, err := h.Service.GetUsers(c.Context(), id)
// 	if err != nil {
// 		if errors.Is(err, comment.ErrFetchingComment) {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	if err := json.NewEncoder(w).Encode(usr); err != nil {
// 		panic(err)
// 	}
// }

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

// PostComment - adds a new comment
// func (h *Handler) PostUser(c *fiber.Ctx) {
// 	var postUserReq PostUserRequest
// 	if err := json.NewDecoder(c.Body).Decode(&postusrReq); err != nil {
// 		return
// 	}

// 	usr := PostUserRequest(postusrReq)
// 	usr, err = h.Service.PostUser(c.Context(), usr)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	if err := json.NewEncoder(w).Encode(usr); err != nil {
// 		panic(err)
// 	}
// }

// UpdateCommentRequest -
type UpdateUserRequest struct {
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

// convert the validated struct into something that the service layer understands
// this is a little verbose, but it allows us to remove tight coupling between our components
func userFromUpdateUserRequest(u UpdateUserRequest) user.User {
	return user.User{
		FullName: u.FullName,
		Age:      u.Age,
		Address:  u.Address,
	}
}

// UpdateComment - updates a comment by ID
// func (h *Handler) UpdateUser(c *fiber.Ctx) {
// 	id := c.Params("id")

// 	var updateUserRequest UpdateUserRequest
// 	if err := json.NewDecoder(r.Body).Decode(&updateUserRequest); err != nil {
// 		return
// 	}

// 	usr := commentFromUpdateCommentRequest(updateUserRequest)

// 	usr, err = h.Service.UpdateComment(r.Context(), commentID, usr)
// 	if err != nil {
// 		log.Error(err.Error())
// 		return
// 	}
// 	if err := json.NewEncoder(w).Encode(usr); err != nil {
// 		panic(err)
// 	}
// }

// GetComment - получить и показать комментарий по ID
func (h *Handler) GetComments(c *fiber.Ctx) error {
	comment, err := h.Service.GetUsers(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(comment)
}

// PostComment - добавить новый комментарий
func (h *Handler) PostComment(c *fiber.Ctx) error {
	var postCmtReq PostUserRequest

	if err := c.BodyParser(&postCmtReq); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	cmt := userFromPostUserRequest(postCmtReq)
	cmt, err := h.Service.PostUser(c, cmt)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(cmt)
}

// UpdateComment - обновить комментарий
func (h *Handler) UpdateComment(c *fiber.Ctx) error {
	commentID := c.Params("id")

	var updateCmtRequest UpdateUserRequest
	if err := c.BodyParser(&updateCmtRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	cmt := userFromUpdateUserRequest(updateCmtRequest)

	cmt, err := h.Service.UpdateUser(c, commentID, cmt)
	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(cmt)
}
