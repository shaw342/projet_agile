package main

import (
	"fmt"

	"github.com/fauna/fauna-go"
	"github.com/gin-gonic/gin"
)

type Project struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Owner string   `json:"owner"`
	Tasks []string `json:"tasks"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

var client *fauna.Client
var clientERR error

func main() {
	r := gin.New()
	r.Use(corsMiddleware())

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
		data := Project{}

		data.Id = "123e4567-e89b-12d3-a456-426614174000"
		data.Name = "my first task"
		data.Owner = "654f3210-feda-4baf-8765-081235432100"

		ctx.JSON(200, data)
	})

	r.GET("/task/get", func(ctx *gin.Context) {
		data := Task{}

		data.Id = "123e4567-e89b-12d3-a456-426614174000"
		data.Name = "create user interface"
		data.State = "INPROGRESS"

		ctx.JSON(200, data)
	})

	r.GET("/user/get", func(ctx *gin.Context) {
		data := User{}

		data.Id = "123e4567-e89b-12d3-a456-426614174000"
		data.Name = "shawan"

		ctx.JSON(200, data)
	})

	r.POST("/user", createCustomer)

	r.POST("/task", createTask)

	r.POST("/project", createProject)
	r.Run()
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func createCustomer(ctx *gin.Context) {
	data := User{}

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(404, ctx.Errors)
		return
	}

	createUser, err := fauna.FQL(`User.create(${data})`, map[string]any{"data": data})

	if err != nil {
		panic(err)
	}
	res, err := client.Query(createUser)
	if err != nil {
		panic(err)
	}
	var scout User

	if err := res.Unmarshal(&scout); err != nil {
		panic(err)
	}

	fmt.Println(scout.Name)
	ctx.JSON(200, scout)
}

func createTask(ctx *gin.Context) {
	task := Task{}

	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(404, ctx.Errors)
		return
	}

	createTask, err := fauna.FQL(`Task.create(${task})`, map[string]any{"task": task})

	if err != nil {
		panic(err)
	}

	res, err := client.Query(createTask)
	if err != nil {
		panic(err)
	}

	var scout Task

	if err := res.Unmarshal(&scout); err != nil {
		panic(err)
	}
	fmt.Println(scout.Name)
	ctx.JSON(200, scout)
}

func createProject(ctx *gin.Context) {
	project := project{}

	if err := ctx.BindJSON(&project); err != nil {
		ctx.JSON(404, ctx.Errors)
		return
	}

	createProject, err = fauna.FQL("Porjects.create(${project})",map[string]any{"project":project})

	if err != nil {
		panic(err)
	}

	res, err = client.Query(createProject)

	if err != nil {
		panic(err)
	}

	var scout Project

	if err := res.Unmarchal(&scout); err != nil {
		panic(err)
	}

	fmt.Println(scout.Name)

	ctx.JSON(200, scout)
}
