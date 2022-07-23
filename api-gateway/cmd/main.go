package main

import (
	"log"

	"github.com/adhtanjung/api-gateway/pkg/asset"
	"github.com/adhtanjung/api-gateway/pkg/auth"
	"github.com/adhtanjung/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")

	authSvc := *auth.RegisterRoutes(v1, &c)
	asset.RegisterRoutes(v1, &c, &authSvc)

	r.Run(c.Port)

}
