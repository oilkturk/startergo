package main

import (
	"log"
	"starter/model"
	"starter/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	initServices()
}

func initServices() {
	connectionString := "user=postgres password=postgres dbname=starterDB sslmode=disable"
	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		log.Print(err)
	}

	db.SingularTable(true)

	db.LogMode(true)

	db.AutoMigrate(&model.Todo{})
	defer db.Close()

	todoService := &service.TodoService{DB: db}

	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	authorized := r.Group("/")
	{
		//user services
		authorized.GET("/starter/todo/add", todoService.Add)
		authorized.GET("/starter/todo/update", todoService.Update)
		authorized.GET("/starter/todo/get", todoService.Get)
		authorized.GET("/starter/todo/delete", todoService.Delete)
	}

	r.Run(":8080")
}
