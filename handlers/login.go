package handlers

import (
	"badminton-app/database"
	"badminton-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"error": ""})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	result := database.DB.Where("username = ? AND password = ?", username, password).First(&user)

	if result.Error != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Username atau password salah"})
		return
	}

	c.SetCookie("user", username, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/login")
}
