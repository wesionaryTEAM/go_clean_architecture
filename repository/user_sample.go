package repository

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users []User
)

func init() {
	users = []User{User{Id: 1, Name: "User1", Email: "user1@gmail.com"}}
}

func GetUsers(c *gin.Context) {
	result := users
	c.JSON(http.StatusOK, result)
}

func AddUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = len(users) + 1
	users = append(users, user)
	c.JSON(http.StatusOK, user)
}
