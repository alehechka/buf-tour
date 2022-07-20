package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/alehechka/buf-tour/petstore/gen/proto/go/pet/v1"
	"github.com/alehechka/buf-tour/petstore/utils"
)

var (
	// gRPC server endpoint
	grpcServerEndpoint = utils.GetEnv("GRCP-SERVER-ENDPOINT", "127.0.0.1:8080")
)

func main() {
	mux, err := createRESTHandler()
	utils.Check(err)

	engine := gin.Default()
	engine.Any("/*path", gin.WrapH(mux))
	utils.Check(engine.Run())
}

func createRESTHandler() (mux *runtime.ServeMux, err error) {
	ctx := context.Background()

	mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err = gw.RegisterPetStoreServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		return nil, err
	}

	return
}
