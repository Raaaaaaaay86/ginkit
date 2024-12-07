package controller

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raaaaaaaay86/ginkit"
	"github.com/raaaaaaaay86/ginkit/example/route/middleware"
)

var _ ginkit.RouteGroup = (*User)(nil)

type User struct {
}

// GetRoutes returns the list of routes that will be registered into gin.Engine
func (u *User) GetRoutes() []ginkit.RouteFactory {
	return []ginkit.RouteFactory{
		u.GetById, // Don't forget to add method in the slice which going to be registered
	}
}

// v1 is a helper function to create a path with /v1 prefix
func (u *User) v1(path string) ginkit.Path {
	return ginkit.Path{
		Name: fmt.Sprintf("/v1%s", path),
		Before: []gin.HandlerFunc{
			middleware.PrintMessage("v1: before"),
			middleware.PrintMessage("authentication..."),
		},
		After: []gin.HandlerFunc{
			middleware.PrintMessage("v1: after"),
		},
	}
}

func (u *User) GetById() ginkit.Route {
	return ginkit.Route{
		Method: http.MethodGet,
		Path:   u.v1("/user/:id"), // This router will be registered as /v1/user/:id
		Before: []gin.HandlerFunc{
			middleware.PrintMessage("GetById: before"),
		},
		After: []gin.HandlerFunc{
			middleware.PrintMessage("GetById: after"),
		},
		Handler: func(ctx *gin.Context) {
			slog.Info("GetById: handler")

			ctx.JSON(http.StatusOK, gin.H{"id": ctx.Param("id"), "name": "john doe"})
		},
	}
}

func NewUser() *User {
	return &User{}
}
