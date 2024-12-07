package ginkit

import "github.com/gin-gonic/gin"

// Path is a struct that contains the path information.
// It's like the `gin.Engine.Group` but expressed in a struct
type Path struct {
	Name   string            // Route name
	Before []gin.HandlerFunc // Middleware before the handler
	After  []gin.HandlerFunc // Middleware after the handler
}

// Route is a struct that contains the route information.
// It's like the `gin.Engine.Handle` but expressed in a struct
type Route struct {
	Method  string            // HTTP method
	Path    Path              // Path information
	Before  []gin.HandlerFunc // Middleware before the handler
	Handler gin.HandlerFunc   // Main Handler
	After   []gin.HandlerFunc // Middleware after the handler
}

// Handlers returns the list of handlers in the correct order.
//
// The execution order is path.before -> route.before -> route.handler -> route.after -> path.after
func (r Route) Handlers() []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, 0)

	handlers = append(handlers, r.Path.Before...)
	handlers = append(handlers, r.Before...)
	handlers = append(handlers, r.Handler)
	handlers = append(handlers, r.After...)
	handlers = append(handlers, r.Path.After...)

	return handlers
}

// RouteFactory is a function that returns a Route
type RouteFactory func() Route

// RouteGroup is a group of routes that can be registered to the gin.Engine
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
