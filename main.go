package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

// 定义数据结构体

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User

//var Users2 = []User{}

func main() {
	r := gin.Default()

	// 设置请求路由信息

	userRoutes := r.Group("/users") // 基于/users 目录为基准目录
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", EditUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}
	// 设置监听端口信息
	if err := r.Run(":6000"); err != nil {
		log.Fatal(err.Error())
	}
}

// api函数

func GetUsers(c *gin.Context) {
	c.JSON(200, Users)

}

func CreateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid request body, create user filed !",
		})
		return
	}
	reqBody.ID = uuid.New().String()
	Users = append(Users, reqBody)
	c.JSON(200, gin.H{
		"error":   false,
		"message": "Create User Successfully !",
	})
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid request body, edit user failed !",
		})
		return
	}
	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age
			c.JSON(200, gin.H{
				"error":   false,
				"message": "Edit user successfully !",
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "No user !",
	})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			c.JSON(200, gin.H{
				"error":   false,
				"message": "Delete user successfully! ",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "Invalid id params",
	})
}
