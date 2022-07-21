package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/adhtanjung/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetAll(ctx *gin.Context, c pb.AuthServiceClient) {

	log.Println("he")
	res, err := c.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
