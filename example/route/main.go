package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raaaaaaaay86/ginutil"
	"github.com/raaaaaaaay86/ginutil/example/route/controller"
	"github.com/raaaaaaaay86/ginutil/example/route/persistence/persistenceimpl"
	"github.com/raaaaaaaay86/ginutil/example/route/service/serviceimpl"
)

func main() {
	if err := Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func Run(ctx context.Context) error {
	engine := gin.Default()

	ginutil.RouteGroups{
		controller.NewUser(),
		controller.NewStore(serviceimpl.NewStore(persistenceimpl.NewStore())),
	}.Register(engine)

	return engine.Run(":8080")
}
