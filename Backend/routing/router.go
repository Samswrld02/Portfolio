package routing

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// router
type Router struct {
	Server *echo.Echo
	port   string
}

func NewRouter() *Router {
	router := &Router{
		Server: echo.New(),
		port:   ":" + os.Getenv("GO_PORT"),
	}

	//set routes for the api
	router.setRoutes()

	return router
}

// setting routes
func (r *Router) setRoutes() {
	api := r.Server.Group("/api", middleware.RequestLogger(), middleware.Recover())

	api.GET("/hi", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"routing": "works"})
	})

	api.GET("/poop", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"poop": "gay"})
	})

	api.GET("/sam", func(c *echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})
}

// starting server
func (r *Router) Start() error {
	return r.Server.Start(r.port)
}
