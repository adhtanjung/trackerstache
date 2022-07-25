package auth

import (
	"github.com/adhtanjung/api-gateway/pkg/auth/routes"
	"github.com/adhtanjung/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.GET("/", svc.GetAll)

	// ROLE
	routes.POST("/role", svc.CreateRole)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) GetAll(ctx *gin.Context) {
	routes.GetAll(ctx, svc.Client)
}

func (svc *ServiceClient) CreateRole(ctx *gin.Context) {
	routes.CreateRole(ctx, svc.Client)
}
