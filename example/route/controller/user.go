package controller

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raaaaaaaay86/ginutil"
	"github.com/raaaaaaaay86/ginutil/example/route/middleware"
)

var _ ginutil.RouteGroup = (*User)(nil)

type User struct {
}

func (u *User) GetRoutes() []ginutil.RouteFactory {
	return []ginutil.RouteFactory{
		u.GetById,
	}
}

func (u *User) v1(path string) ginutil.Path {
	return ginutil.Path{
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

func (u *User) GetById() ginutil.Route {
	return ginutil.Route{
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
