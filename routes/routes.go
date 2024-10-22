package routes

import (
	"net/http"
	"webapp/controllers"
	"webapp/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", middleware.AuthRequired, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
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
