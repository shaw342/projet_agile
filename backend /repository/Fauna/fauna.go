package repository

import (
	"fmt"

	"github.com/fauna/fauna-go"
	"github.com/gin-gonic/gin"
	"github.com/shaw342/projet_argile/backend/model"
)


func newFaunaClient() *fauna.Client {
	client,err := fauna.NewDefaultClient()
	if err != nil{
		panic(err)
	}
	return client
}

func CreateCustomer(ctx *gin.Context) {
	client := newFaunaClient()
	data := model.User{}

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
	var scout model.User

	if err := res.Unmarshal(&scout); err != nil {
		panic(err)
	}

	fmt.Println(scout.Name)
	ctx.JSON(200, scout)
}


func CreateTask(ctx *gin.Context) {
	client := newFaunaClient()
	task := model.Task{}

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

	var scout model.Task

	if err := res.Unmarshal(&scout); err != nil {
		panic(err)
	}
	fmt.Println(scout.Name)
	ctx.JSON(200, scout)
}

func CreateProject(ctx *gin.Context) {
	client := newFaunaClient()
	project := model.Project{}

	if err := ctx.BindJSON(&project); err != nil {
		ctx.JSON(404, ctx.Errors)
		return
	}

	createProject, err := fauna.FQL("Projects.create(${project})",map[string]any{"project":project})

	if err != nil {
		panic(err)
	}

	res, err := client.Query(createProject)

	if err != nil {
		panic(err)
	}

	var scout model.Project

	if err := res.Unmarshal(&scout); err != nil {
		panic(err)
	}

	fmt.Println(scout.Name)

	ctx.JSON(200, scout)
}
