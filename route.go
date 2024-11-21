package ginutil

import "github.com/gin-gonic/gin"

type Path struct {
	Name   string
	Before []gin.HandlerFunc
	After  []gin.HandlerFunc
}

type Route struct {
	Method  string
	Path    Path
	Before  []gin.HandlerFunc
	Handler gin.HandlerFunc
	After   []gin.HandlerFunc
}

func (r Route) Handlers() []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, 0)

	handlers = append(handlers, r.Path.Before...)
	handlers = append(handlers, r.Before...)
	handlers = append(handlers, r.Handler)
	handlers = append(handlers, r.After...)
	handlers = append(handlers, r.Path.After...)

	return handlers
}

type RouteFactory func() Route

type RouteGroup interface {
	GetRoutes() []RouteFactory
}

type RouteGroups []RouteGroup

func (rg RouteGroups) Register(engine *gin.Engine) {
	for _, r := range rg {
		for _, getRouter := range r.GetRoutes() {
			router := getRouter()
			engine.Handle(router.Method, router.Path.Name, router.Handlers()...)
		}
	}
}
