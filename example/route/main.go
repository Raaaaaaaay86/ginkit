package main

import (
	"context"
	"log"
	"net/http"

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
	server := BuildServer()
	opt := ginutil.ServeOptions{}

	if err := ginutil.Serve(ctx, server, opt); err != nil {
		return err
	}

	return nil
}

func BuildServer() *http.Server {
	engine := gin.Default()

	ginutil.RouteGroups{
		controller.NewUser(),
		controller.NewStore(serviceimpl.NewStore(persistenceimpl.NewStore())),
	}.Register(engine)

	return &http.Server{
		Handler: engine.Handler(),
	}
}
