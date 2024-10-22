package routes

import (
	"net/http"
	"webapp/controllers"
	"webapp/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	// Dashboard route with authentication middleware + pass user data to the template
	router.GET("/", middleware.AuthRequired, func(c *gin.Context) {
		user := c.MustGet("user").(string)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"user": user,
		})
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	router.POST("/register", controllers.Register)

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", controllers.Login)

	router.GET("/logout", controllers.Logout)
}
