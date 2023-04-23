package entity

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UsersLogin struct {
	UserEmail string `bson:"user_email" json:"user_email"`
	Password  string `bson:"password" json:"password,omitempty"`
}

type Users struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserName     string             `bson:"user_name" json:"user_name"`
	UserLastName string             `bson:"user_last_name" json:"user_last_name"`
	UserEmail    string             `bson:"user_email" json:"user_email"`
	Password     string             `bson:"password" json:"password,omitempty"`
	CreatedAt    string             `bson:"created_at" json:"created_at"`
	UpdatedAt    string             `bson:"updated_at" json:"updated_at"`
	Permissions  Permissions        `bson:"permissions" json:"permissions"`
	UserRole     bool               `bson:"role" json:"role"`
}

type RolesPermissions struct {
	PermissionCreate []string
	PermissionRead   []string
	ReadRegister     []string
	Roles            []string
}

type Permissions struct {
	CreateTopic bool `bson:"create_topic" json:"create_topic"`
	ReadTopic   bool `bson:"read_topic" json:"read_topic"`
}

func (u Users) UsersValidations(ctx *fiber.Ctx) error {

	if len(u.UserEmail) == 0 {
		return ctx.JSON(fiber.Map{"user_email": "User Email is required ..."})
	}
	if len(u.Password) < 8 {
		return ctx.JSON(fiber.Map{"password": "Password must be at least lenght of 6 elements ..."})
	}
	return nil
}

func SetPassword(password string) (string, error) {
	cost := 8
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Panic("Error ")
	}
	return string(hashedPassword), nil
}

func ComparePassword(s string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(s), []byte(password))
}
