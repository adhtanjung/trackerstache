package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/adhtanjung/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RegisterRequestBody struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	WorkDate time.Time `json:"work_date"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Username: body.Username,
		Password: body.Password,
		WorkDate: timestamppb.New(body.WorkDate),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
