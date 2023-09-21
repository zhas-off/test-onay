package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/zhas-off/test-onay/internal/user"
)

type UserRow struct {
	ID       string
	FullName string
	Age      int
	Address  string
}

// GetUsers - retrieves a users from the database by ID
func (d *Database) GetUsers(ctx *fiber.Ctx) (user.User, error) {
	// fetch CommentRow from the database and then convert to comment.Comment
	var users UserRow
	err := d.Client.Find(&users).Error
	if err != nil {
		return user.User{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	// sqlx with context to ensure context cancelation is recognized
	return user.User(users), nil
}

// PostComment - adds a new comment to the database
func (d *Database) PostUser(ctx *fiber.Ctx, usr user.User) (user.User, error) {
	usr.ID = uuid.NewV4().String()
	postRow := UserRow{
		ID:       usr.ID,
		FullName: usr.FullName,
		Age:      usr.Age,
		Address:  usr.Address,
	}

	err := d.Client.Create(&postRow).Error
	if err != nil {
		return user.User{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	return usr, nil
}

// UpdateComment - updates a comment in the database
func (d *Database) UpdateUser(ctx *fiber.Ctx, uuid string, usr user.User) (user.User, error) {
	postRow := UserRow{
		ID:       uuid,
		FullName: usr.FullName,
		Age:      usr.Age,
		Address:  usr.Address,
	}

	err := d.Client.Save(&postRow).Error
	
	if err != nil {
		return user.User{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	return usr, nil
}
