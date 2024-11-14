package main

import (
	"go_project/src/controller"
	"go_project/src/middleware"
	"go_project/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// je suis un test
func setupRouter() *gin.Engine {

	var loginService service.LoginService = service.NewLoginService()
	var jwtService service.JWTService = service.NewJWTService()
	var loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// Ping test
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "OK"})

	})

	apiRoutes := r.Group("/api", middleware.AuthorizeJWT())
	{
		apiRoutes.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
