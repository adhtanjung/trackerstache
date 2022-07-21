package routes

import (
	"context"
	"net/http"

	"github.com/adhtanjung/api-gateway/pkg/asset/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateAssetRequestBody struct {
	Name       string                 `json:"name"`
	Stock      int64                  `json:"stock"`
	Price      int64                  `json:"price"`
	BoughtDate *timestamppb.Timestamp `json:"bought_date"`
}

func CreateAsset(ctx *gin.Context, c pb.AssetServiceClient) {
	body := CreateAssetRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateAsset(context.Background(), &pb.CreateAssetRequest{
		Name:       body.Name,
		Stock:      body.Stock,
		Price:      body.Price,
		BoughtDate: body.BoughtDate,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
