package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"publisher_service/config"
	"publisher_service/internal/domain/entity"
	"publisher_service/internal/infrastructure/datasource/mongodb"
	"publisher_service/internal/middleware"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessService struct {
	db       mongodb.MongoService
	security middleware.SecurityService
	conf     config.Config
}

func (c *AccessService) compareCredentials(ctx *fiber.Ctx, access entity.UsersLogin, user string, password string) (bool, error) {

	if err := entity.ComparePassword(password, access.Password); err != nil {
		return false, ctx.JSON(fiber.Map{"error": "Invalid_access"})
	}
	if user != access.UserEmail {
		return false, ctx.JSON(fiber.Map{"error": "Invalid_access"})
	}
	fmt.Print(password)
	fmt.Print(access.Password)
	return true, nil
}

func (c *AccessService) RegisterUser(ctx context.Context, requestChan <-chan entity.Users) error {

	dataBase, collection, _ := c.conf.MongoDBConfig()

	select {
	case requests := <-requestChan:
		mongoCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		conn := c.db.ConnMongoDB().Database(dataBase)
		coll := conn.Collection(collection)

		adminUser := entity.Users{
			UserName:     requests.UserName,
			UserLastName: requests.UserLastName,
			UserEmail:    requests.UserEmail,
			Password:     requests.Password,
			CreatedAt:    time.Now().Local().String(),
			UserRole:     requests.UserRole,
			ID:           primitive.NewObjectID(),
			Permissions: entity.Permissions{
				CreateTopic: requests.Permissions.CreateTopic,
				ReadTopic:   requests.Permissions.ReadTopic,
			},
		}

		hash, err := entity.SetPassword(requests.Password)
		if err != nil {
			log.Println(err)
		}
		adminUser.Password = string(hash)

		_, err = coll.InsertOne(mongoCtx, adminUser)
		if err != nil {
			return fmt.Errorf("failed to execute query: %v", err)
		}

		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *AccessService) UserLogin(ctx *fiber.Ctx, requestChan <-chan entity.UsersLogin) string {

	request := <-requestChan
	findChan := c.FindUserByEmail(request.UserEmail)
	user := <-findChan

	t, err := c.compareCredentials(ctx, request, user.UserEmail, user.Password)
	if err != nil {
		log.Println(err)
	}

	isSuperAdmin := strings.Contains(ctx.Path(), "/service/publisher")
	var scope string
	if isSuperAdmin {
		scope = "user"
	} else {
		scope = "super_admin"
	}
	per := RolesPermissionsAssignment(user)

	if t {

		nowTime := time.Now()
		expireTime := nowTime.Add(12 * time.Hour)
		token, errL := c.security.GenerateJWT(user.UserEmail, scope, []string{"user"}, per)
		if errL != nil {
			ctx.Status(http.StatusOK).JSON(&fiber.Map{
				"message": "Invalid Credentials ..."})
		}

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  expireTime,
			HTTPOnly: true,
		}

		ctx.Cookie(&cookie)
		return "{access: successed}"

	} else {
		return "{error: invalid}"
	}
}

func (c *AccessService) FindUserByEmail(email string) chan entity.Users {
	dataBase, collection, _ := c.conf.MongoDBConfig()
	userChan := make(chan entity.Users)
	go func() {
		var adminUser entity.Users
		condition := bson.M{
			"user_email": email,
		}
		mongoCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		conn := c.db.ConnMongoDB().Database(dataBase)
		coll := conn.Collection(collection)
		err := coll.FindOne(mongoCtx, condition).Decode(&adminUser)
		if err != nil {
			log.Println("No User found ...", err.Error())
			userChan <- entity.Users{} // send empty model to channel
			return
		}

		userChan <- adminUser // send model to channel
	}()

	return userChan
}

func RolesPermissionsAssignment(user entity.Users) []string {
	permissionsMap := map[string]string{
		"CreateTopic": "create_topic",
		"ReadTopic":   "read_topic",
	}

	inventoryPer := []string{}

	if user.Permissions.CreateTopic {
		inventoryPer = append(inventoryPer, permissionsMap["CreateTopic"])
	}
	if user.Permissions.ReadTopic {
		inventoryPer = append(inventoryPer, permissionsMap["ReadTopic"])
	}

	return inventoryPer
}
