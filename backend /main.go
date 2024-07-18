package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/shaw342/projet_argile/backend/Verification"
	"github.com/shaw342/projet_argile/backend/repository/Fauna"
)


func main() {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"hello": "world",
		})
	})

	v1 := r.Group("api/v1")
	{
		v1.POST("/user", repository.CreateUser)
		v1.POST("/task", repository.CreateTask)
		v1.POST("/project", repository.CreateProject)
		v1.DELETE("/deleteProject",repository.DeleteProject)
		v1.DELETE("/deleteTask",repository.DeleteTask)
		v1.PATCH("/updateProject",repository.UpdateProject)
		v1.PATCH("/updateTask",repository.UpdateTasks)
		v1.GET("/task/get",repository.GetTask)
		v1.GET("/project/get",repository.GetProject)
		v1.GET("/user/get",repository.GetUser)

	}
	r.Run()
}
