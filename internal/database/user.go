package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/zhas-off/test-onay/internal/user"
)

// UserRow - как наши пользователи выглядят в базе данных
type UserRow struct {
	ID       string
	FullName string
	Age      int
	Address  string
}

// GetUsers - извлекает всех пользователей из базы данных
func (d *Database) GetUsers(ctx *fiber.Ctx) (user.User, error) {
	var users UserRow
	err := d.Client.Find(&users).Error
	if err != nil {
		return user.User{}, fmt.Errorf("failed to insert user: %w", err)
	}

	return user.User(users), nil
}

// PostUser - добавляет нового пользователя в базу данных
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
		return user.User{}, fmt.Errorf("failed to insert user: %w", err)
	}

	return usr, nil
}

// UpdateUser - обновляет пользователя в базе данных
func (d *Database) UpdateUser(ctx *fiber.Ctx, uuid string, usr user.User) (user.User, error) {
	postRow := UserRow{
		ID:       uuid,
		FullName: usr.FullName,
		Age:      usr.Age,
		Address:  usr.Address,
	}

	err := d.Client.Save(&postRow).Error
	
	if err != nil {
		return user.User{}, fmt.Errorf("failed to insert user: %w", err)
	}

	return usr, nil
}
