package main

import (
	"alexnerotd/blog/handlers"
	"alexnerotd/blog/models"
	"alexnerotd/blog/repositories"
	"alexnerotd/blog/services"
	"alexnerotd/blog/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/contrib/cors"
)

func main() {

	db, err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Blog{})
	if err != nil {
		panic("Error migrating Blog table")
	}

	useRepository := repositories.NewBlogRepository(db)
	useService := services.NewBlogService(useRepository)
	useHandler := handlers.NewBlogHandler(useService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
	}))

	router.GET("/blog/:id", useHandler.GetBlogById)
	router.POST("/blog", useHandler.CreateBlog)

	router.Run(":8080")
}
