package interfaces

import (
	"publisher_service/config"
	"publisher_service/internal/domain/entity"
	"publisher_service/internal/infrastructure/websocket"
	"publisher_service/internal/middleware"
	metrics "publisher_service/pkg/monitoring/middleware"

	"github.com/gofiber/fiber/v2"
)

type UsersRouterService struct {
	conf config.Config
}

func (c UsersRouterService) PublisherRouter(app *fiber.App) {

	security := middleware.SecurityService{}

	roles, permCreate, permRead := c.conf.RolesPermissions()
	rp := entity.RolesPermissions{
		Roles:            roles,
		PermissionCreate: permCreate,
		PermissionRead:   permRead,
	}
	hand := UsersHandlerService{}
	access := HandlerService{}

	s := app.Group("service")
	p := s.Group("publisher", metrics.RecordRequestLatency)

	p.Post("register/user", access.RegisterAdminUserHandler)
	p.Post("login", access.AdminUserLoginHandler)
	p.Post("api/send", websocket.WebSocketHandler)

	adminAuthenticated := p.Use(security.IsAuthenticated)

	permissionCreate := adminAuthenticated.Use(security.ValidateRolesAndPermissions(rp.Roles, rp.PermissionCreate))

	permissionCreate.Post("create/topic", hand.CreateTopicHandler)

}
