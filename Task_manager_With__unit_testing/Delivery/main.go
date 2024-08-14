package main

import (
	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"

	//infrastructure "task-manager/Infrastructure"
	repositories "task-manager/Repositories"
	usecase "task-manager/Usecases"

	//"time"
	"log"

	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	//client := infrastructure.MongoDBInit() //mongodb initialization
	mongoURI := "mongodb+srv://<username>:<password>.@cluster0.fek5tj1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	//initialization of the repositories
	task_repo := repositories.NewTaskRepository(client)
	user_repo := repositories.NewUserRepository(client)

	//set-up the controllers
	cont := controllers.NewTaskController(usecase.NewTaskUsecase(task_repo), usecase.NewUserUsecase(user_repo))

	//the router gateway
	router := gin.Default()
	routers.SetupRoutes(router, cont)
	router.Run(":8080")
}
