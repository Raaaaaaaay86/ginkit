package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raaaaaaaay86/ginutil"
	"github.com/raaaaaaaay86/ginutil/example/route/controller"
	"github.com/raaaaaaaay86/ginutil/example/route/persistence/persistence_impl"
	"github.com/raaaaaaaay86/ginutil/example/route/service/service_impl"
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
		controller.NewStore(service_impl.NewStore(persistence_impl.NewStore())),
	}.Register(engine)

	return engine.Run(":8080")
}
