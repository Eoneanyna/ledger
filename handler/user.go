package handler

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func RegisterUser(r *gin.Context) {

}

func LoginUser(r *gin.Context) {

}
