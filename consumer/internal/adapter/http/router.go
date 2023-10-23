package httpHandler

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
}

type Credentials struct {
	Host string
	Port string
}

// NewRouter creates a new HTTP router
func NewRouter(orderHandler *OrderHandler) (*Router, error) {
	// gin.SetMode(gin.ReleaseMode)

	// config := cors.DefaultConfig()
	// allowedOrigins := os.Getenv("HTTP_ALLOWED_ORIGINS")
	// originsList := strings.Split(allowedOrigins, ",")
	// config.AllowOrigins = originsList

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		user := v1.Group("/order")
		{
			user.GET("/:uid", orderHandler.GetOrderByUId)
			//user.GET("/list", orderHandler.GetListOrders)
		}
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(cr *Credentials) error {
	listenAddr := fmt.Sprintf("%s:%s", cr.Host, cr.Port)

	return r.Run(listenAddr)
}

// validationError sends an error response for some specific request validation error
func validationError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, err)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func handleError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, err)
}
