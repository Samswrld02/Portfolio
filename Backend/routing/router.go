package routing

import (
	"os"

	"github.com/Samswrld02/Portfolio/auth"
	"github.com/Samswrld02/Portfolio/controllers"
	"github.com/Samswrld02/Portfolio/middlewareC"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	middleware "github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

// router
type Router struct {
	Server *echo.Echo
	port   string
	db     *gorm.DB
	auth   *auth.Auth
}

type CustomValidator struct {
	validator *validator.Validate
}

func (v *CustomValidator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func NewRouter(db *gorm.DB) *Router {
	router := &Router{
		Server: echo.New(),
		port:   ":" + os.Getenv("GO_PORT"),
		db:     db,
		auth:   &auth.Auth{Db: db},
	}

	//set validatopr
	router.setValidator()

	//set routes for the api
	router.setRoutes()

	return router
}

// setting routes
func (r *Router) setRoutes() {
	api := r.Server.Group("/api", middleware.RequestLogger(), middleware.Recover())

	api.POST("/login", r.auth.Login)
	api.POST("/register", r.auth.Register)

	//project crud, inject db into controller
	projectController := controllers.NewProjectController(r.db)
	project := api.Group("/projects")
	project.GET("", projectController.Read)
	project.GET("/:id", projectController.Show)

	protectedProjects := project.Group("", middlewareC.CheckJwt)
	protectedProjects.POST("", projectController.Create)
	protectedProjects.PATCH("/:id", projectController.Update)
	protectedProjects.DELETE("/:id", projectController.Delete)
}

func (r *Router) setValidator() {
	r.Server.Validator = &CustomValidator{validator: validator.New()}
}

// starting server
func (r *Router) Start() error {
	return r.Server.Start(r.port)
}
