package main

import (
	"fmt"
	"net/http"
	"sqlite"

	"github.com/gin-gonic/gin"
)

func main() {

	/****************
	* Connect to DB *
	*****************/

	db, err := sqlite.ConnectDB(sqlite.Test_db_path)
	if err != nil {
		fmt.Println("Error when trying to connect to db")
		return
	}
	defer db.Close()

	/****************
	* Setup Gin API *
	*****************/

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	/********************
	* Authenticate User *
	*********************/

	// POST logic for login
	r.POST("/login", func(c *gin.Context) {
		// Struct for user data, includes username and password
		var loginData struct {
			Username string `json:"Username"`
			Password string `json:"Password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		// Authentification logic for user
		// 1. Check if user exists in the db
		user, err := sqlite.GetUser(db, loginData.Username)
		if err != nil || user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		} else {
			// 2. Check if the password is correct
			if sqlite.VerifyUser(db, loginData.Username, loginData.Password) {
				c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			}
		}

	})

	// ✅ Start server on port 8080
	r.Run(":8080")
}

// To run go code: go run main.go
// To build go code: go build -o myapp
