package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

	router := gin.Default()

	// Endpoint untuk pendaftaran
	router.POST("/register", registerHandler)

	// Menjalankan server pada port 8080
	router.Run(":8080")
}

func registerHandler(c *gin.Context) {
	var newUser User

	// Bind JSON ke struct User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lakukan validasi atau operasi lainnya sesuai kebutuhan aplikasi Anda
	// ...

	// Tampilkan respons berhasil jika semuanya berjalan dengan baik
	c.JSON(http.StatusOK, gin.H{"message": "Pendaftaran berhasil", "user": newUser})
}
