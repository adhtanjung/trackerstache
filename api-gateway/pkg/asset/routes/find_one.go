package routes

import (
	"context"
	"net/http"

	"github.com/adhtanjung/api-gateway/pkg/asset/pb"
	"github.com/gin-gonic/gin"
)

func FineOne(ctx *gin.Context, c pb.AssetServiceClient) {
	id := ctx.Param("id")

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
