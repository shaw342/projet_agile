package repository

import (
	"fmt"
	"net/http"

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


func CreateUser(ctx *gin.Context) {
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

func UpdatePassword(ctx *gin.Context){
	client := newFaunaClient()
	user := model.User{}
	if err := ctx.ShouldBindJSON(user); err != nil{
		ctx.JSON(404,err)
	}
	id := GetId(user.Name,client)
	command := fmt.Sprintf(`User.byId("%s")?.update({"Password": "%s"})`,id)
	query, err := fauna.FQL(command,nil)
	if err != nil{
		panic(err)
	}
	res,err := client.Query(query)

	if err != nil{
		panic(err)
	}

	var result model.User

	if err := res.Unmarshal(&result); err != nil{
		panic(err)
	}

	ctx.JSON(200,result.Password)
}


func GetProjectId(name string,client *fauna.Client) string{
	var Id string
	query,err := fauna.FQL("Pojects.byName(${name})",map[string]any{"name":name})

	if err != nil{
		panic(err)
	}
	res,_ := client.Query(query)

	if err := res.Unmarshal(&Id); err != nil{
		panic(err)
	}

	return Id
}

func GetByName(name string, client *fauna.Client) model.Project{
	var result model.Project

	query,err := fauna.FQL(`Project.byName(${name})`,map[string]any{"name":name})

	if err != nil{
		panic(err)
	}

	res,_ := client.Query(query)

	if err := res.Unmarshal(&result); err != nil{
		panic(err)
	}

	return result
}

func DeletProject(ctx *gin.Context){
	client := newFaunaClient()
	projectName := model.Project{}
	if err := ctx.ShouldBindJSON(&projectName);err != nil{
		panic(err)
	}
	project := GetByName(projectName.Name,client)

	query,err := fauna.FQL(`${project}!.delete()`,map[string]any{"project":project})
	if err != nil{
		panic(err)
	}
	res,err := client.Query(query)

	if err != nil{
		panic(err)
	}

	ctx.JSON(http.StatusAccepted,res)
}
