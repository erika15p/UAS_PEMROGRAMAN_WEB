package router

import (
	"badminton-app/handlers"
	"badminton-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.Use(func(c *gin.Context) {
		user, err := c.Cookie("user")
		if err == nil {
			c.Set("user", user)
		}
		c.Next()
	})

	r.GET("/login", handlers.ShowLogin)
	r.POST("/login", handlers.Login)
	r.GET("/logout", handlers.Logout)

	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/", handlers.ShowHome)
		auth.GET("/kehadiran", handlers.ListKehadiran)
		auth.POST("/kehadiran", handlers.AddKehadiran)
		auth.POST("/kehadiran/edit/:id", handlers.UpdateKehadiran)
		auth.GET("/kehadiran/delete/:id", handlers.DeleteKehadiran)

		auth.GET("/keuangan", handlers.ListKeuangan)
		auth.POST("/keuangan", handlers.AddKeuangan)
		auth.POST("/keuangan/edit/:id", handlers.UpdateKeuangan)
		auth.GET("/keuangan/delete/:id", handlers.DeleteKeuangan)

		auth.GET("/laporan", handlers.ShowLaporan)
		auth.POST("/laporan", handlers.GenerateLaporan)
		auth.GET("/laporan/export", handlers.ExportPDF)
	}

	return r
}
