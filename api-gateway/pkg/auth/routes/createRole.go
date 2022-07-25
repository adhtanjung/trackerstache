package routes

import (
	"context"
	"net/http"

	"github.com/adhtanjung/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type CreateRoleRequestBody struct {
	Name string `json:"name"`
}

func CreateRole(ctx *gin.Context, c pb.AuthServiceClient) {
	body := CreateRoleRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateRole(context.Background(), &pb.CreateRoleRequest{
		Name: body.Name,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)

}
