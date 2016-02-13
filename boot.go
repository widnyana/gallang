package gallang

import (
	"github.com/gin-gonic/gin"
)

// Gallang provide interface of gallang oembed app
type Gallang struct {
	App  *gin.Engine
	host string
}

// New build a new application
func New(addr string) *Gallang {
	if len(addr) == 0 {
		addr = "0.0.0.0:5000"
	}

	app := new(Gallang)
	app.host = addr
	app.Build()

	return app
}

// Build will do action of building an app
func (g *Gallang) Build() {
	g.App = gin.New()
	g.App.Use(gin.Logger())
	g.App.Use(gin.Recovery())

	routes := getRouteDefinition()

	for _, item := range routes {
		g.RouteHandle(item.method, item.endpoint, item.worker)
	}
}

// Use will tell Gallang to use the middleware
func (g *Gallang) Use(middleware gin.HandlerFunc) {
	g.App.Use(middleware)
}

// RouteHandle provide a route registration
func (g *Gallang) RouteHandle(httpMethod, relativePath string, handlers gin.HandlerFunc) gin.IRoutes {
	return g.App.Handle(httpMethod, relativePath, handlers)
}

// Run gallang API
func (g *Gallang) Run() (err error) {
	err = g.App.Run(g.host)
	return err
}
