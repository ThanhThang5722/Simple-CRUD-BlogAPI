package main

import (
	database "BlogAPI/pkg/database"
	"BlogAPI/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.ConnectMySQL()
	if err != nil {
		fmt.Println("Failt to connect MySQL")
		return
	}
	defer db.Close()
	defer fmt.Println(os.Getenv("test"))

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Ping successful",
			})
		})
	}
	//ROUTER DEFINE
	routes.ArticleRouter(api)
	router.Run()
}
