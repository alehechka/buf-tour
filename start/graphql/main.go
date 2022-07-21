package main

import (
	petv1 "github.com/alehechka/buf-tour/petstore/gen/proto/go/pet/v1"
	"github.com/alehechka/buf-tour/petstore/utils"
	"github.com/friendsofgo/graphiql"
	"github.com/gin-gonic/gin"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func main() {
	mux, err := createGraphQLHandler()
	utils.Check(err)

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	utils.Check(err)

	engine := gin.Default()
	engine.POST("/graphql", gin.WrapH(mux))
	engine.GET("/graphiql", gin.WrapH(graphiqlHandler))
	utils.Check(engine.Run())
}

func createGraphQLHandler() (mux *runtime.ServeMux, err error) {
	mux = runtime.NewServeMux()

	if err = petv1.RegisterPetStoreServiceGraphql(mux); err != nil {
		return nil, err
	}

	return
}
