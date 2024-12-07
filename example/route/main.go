package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ginkit "github.com/raaaaaaaay86/ginkit"
	"github.com/raaaaaaaay86/ginkit/example/route/controller"
	"github.com/raaaaaaaay86/ginkit/example/route/persistence/persistenceimpl"
	"github.com/raaaaaaaay86/ginkit/example/route/service/serviceimpl"
)

func main() {
	if err := Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func Run(ctx context.Context) error {
	server := BuildServer()
	opt := ginkit.ServeOptions{}

	if err := ginkit.Serve(ctx, server, opt); err != nil {
		return err
	}

	return nil
}

func BuildServer() *http.Server {
	engine := gin.Default()

	ginkit.RouteGroups{
		controller.NewUser(),
		controller.NewStore(serviceimpl.NewStore(persistenceimpl.NewStore())),
	}.Register(engine)

	return &http.Server{
		Handler: engine.Handler(),
	}
}
