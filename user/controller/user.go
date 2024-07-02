package controller

import (
	"github.com/gin-gonic/gin"
	"musicRecommendation/user/model"
	"musicRecommendation/user/service"
	"net/http"
)

// TODO: create a base controller inside core

func AddUser(c *gin.Context) {
	var err error
	var user model.User
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.AddUser(user)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{"message": "User added successfully"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

// TODO: add error codes
func AddFriend(c *gin.Context) {
	userName := c.Param("userName")
	var request struct {
		Friend string `json:"friend"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddFriend(userName, request.Friend); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Friend added successfully"})
}
