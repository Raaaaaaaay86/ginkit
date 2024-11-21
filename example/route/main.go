package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raaaaaaaay86/ginutil"
	"github.com/raaaaaaaay86/ginutil/example/route/controller"
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
	}.Register(engine)

	return engine.Run(":8080")
}
