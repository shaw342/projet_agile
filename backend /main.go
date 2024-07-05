package main

import (
	"fmt"


	"github.com/fauna/fauna-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shaw342/projet_argile/backend/Verification"
	"github.com/shaw342/projet_argile/backend/model"
	"github.com/shaw342/projet_argile/backend/repository/Fauna"
	//"github.com/shaw342/projet_argile/backend/repository"
)


var client *fauna.Client
var clientERR error

func main() {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	client, clientERR = fauna.NewDefaultClient()

	if clientERR != nil {
		panic(clientERR)
	}
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"hello": "world",
		})
	})

	r.GET("/project/get", func(ctx *gin.Context) {
		data := model.Project{}

		data.Id = "123e4567-e89b-12d3-a456-426614174000"
		data.Name = "my first task"
		data.Owner = "654f3210-feda-4baf-8765-081235432100"

		ctx.JSON(200, data)
	})

	r.GET("/task/get", func(ctx *gin.Context) {
		data := model.Task{}

		data.Id = "123e4567-e89b-12d3-a456-426614174000"
		data.Name = "create user interface"
		data.State = "INPROGRESS"

		ctx.JSON(200, data)
	})

	r.GET("/user/get", func(ctx *gin.Context) {
		data := model.User{}

		data.Id = "123e4567-e89b-12d3-a456-426614174000"
		data.Name = "shawan"

		ctx.JSON(200, data)
	})

	v1 := r.Group("/v1")
	{
	v1.POST("/user", repository.CreateUser)
	v1.POST("/task", repository.CreateTask)
	v1.POST("/project", repository.CreateProject)
	v1.POST("/changePassword", repository.UpdatePassword)
	v1.POST("/email",mailVerification)
	v1.POST("/delete",repository.DeletProject)
	}
	r.Run()
}


func mailVerification(ctx *gin.Context){
	email := model.User{}

	if err := ctx.ShouldBindJSON(&email);err != nil{
		ctx.JSON(404,err)
	}
	fmt.Print(email.Email)
	verification.SendMail(email.Email)
}