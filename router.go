package gallang

import (
	"github.com/gin-gonic/gin"
	"github.com/widnyana/gallang/handler"
)

// GallangRoute is a generic route structure
type GallangRoute struct {
	method   string
	endpoint string
	worker   gin.HandlerFunc
}

// getRouteDefinition return a slice of available router
// you must manualy add your router definition here.
func getRouteDefinition() (routes []GallangRoute) {
	routes = []GallangRoute{
		GallangRoute{
			method:   "GET",
			endpoint: "/",
			worker:   handler.Hello,
		},
		GallangRoute{
			method:   "GET",
			endpoint: "embed",
			worker:   handler.OEmbedService,
		},
	}

	return routes
}
