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

func (u *User) GetRoutes() []ginkit.RouteFactory {
	return []ginkit.RouteFactory{
		u.GetById,
	}
}

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
		Path:   u.v1("/user/:id"),
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
