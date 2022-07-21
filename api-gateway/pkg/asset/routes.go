package asset

import (
	"github.com/adhtanjung/api-gateway/pkg/asset/routes"
	"github.com/adhtanjung/api-gateway/pkg/auth"
	"github.com/adhtanjung/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/asset")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateAsset)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateAsset(ctx *gin.Context) {
	routes.CreateAsset(ctx, svc.Client)
}
