package repository

import (
	"fmt"

	//"net/http"

	"github.com/fauna/fauna-go"
	"github.com/gin-gonic/gin"
	"github.com/shaw342/projet_argile/backend/model"
)



func NewFaunaClient() *fauna.Client {
	client,err := fauna.NewDefaultClient()
	if err != nil{
		panic(err)
	}
	return client
}


func CreateUser(ctx *gin.Context) {
	client := NewFaunaClient()
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
	client := NewFaunaClient()
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
	client := NewFaunaClient()
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

func GetId(name string, client *fauna.Client) string{
	var Id string
	query,err := fauna.FQL("User.byName(${name}).map(.id).first()",map[string]any{"name":name})
	if err != nil{
		panic(err)
	}
	res,_ := client.Query(query)

	if err := res.Unmarshal(&Id); err != nil{
		panic(err)
	}

	return Id
}

func DeleteProject(ctx *gin.Context){
	client := NewFaunaClient()
	name := model.Project{}
	if err := ctx.ShouldBindJSON(&name);err != nil{
		panic(err)
	}
	delete := fmt.Sprintf(`Projects.byName(%s).first()!.delete()`,name.Name)
	query,_ := fauna.FQL(delete,nil)

	res,_ := client.Query(query)

	ctx.JSON(200,res)
	
}

func DeleteTask(ctx *gin.Context){
	client := NewFaunaClient()
	name := model.Task{}
	if err := ctx.ShouldBindJSON(&name);err != nil{
		panic(err)
	}
	delete := fmt.Sprintf(`Task.byName(%s).first()!.delete()`,name.Name)
	query,_ := fauna.FQL(delete,nil)
	res,_ := client.Query(query)
	ctx.JSON(200,res)
}

func UpdateProject(ctx *gin.Context){
	client := NewFaunaClient()
	updateProject := model.Project{}
	name := model.Project{}
	if err := ctx.ShouldBindJSON(&updateProject); err != nil{
		panic(err)
	}
	if err := ctx.ShouldBindJSON(&name); err != nil{
		panic(err)
	}

	data := fmt.Sprintf(`Projects.byName("%s").first()!.update(${updateProject})`,name.Name,updateProject)
	query := fauna.FQL(data,nil)
	res,err := client.Query(query)
	if err != nil{
		panic(err)
	}

	var newProject model.Project

	if err := res.Unmarshal(&newProject); err != nil{
		panic(err)
	}


}