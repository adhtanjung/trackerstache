package asset

import (
	"fmt"

	"github.com/adhtanjung/api-gateway/pkg/asset/pb"
	"github.com/adhtanjung/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AssetServiceClient
}

func InitServiceClient(c *config.Config) pb.AssetServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AssetSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAssetServiceClient(cc)
}
