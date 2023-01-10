package main

import (
	"github.com/blessedmadukoma/todo-vue-go/database"
	"github.com/blessedmadukoma/todo-vue-go/router"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	app := gin.Default()

	app.Use(cors.Default()) // TODO: Configure this later to lock it down better

	// better/more secure CORS setting
	cfg := cors.DefaultConfig()
	// cfg.AllowOrigins = []string{"*"}
	cfg.AllowOrigins = []string{"http://127.0.0.1:5173"}
	app.Use(cors.Default()) // The better/more secure CORS setting

	app.Use(cors.New(cfg))

	app.GET("/health", func(c *gin.Context) {
		var response struct {
			Health string
		}
		response.Health = "I am healthy! 🚀"
		c.JSON(http.StatusOK, response)
	})

	router.SetupRouter(app)

	app.Run(":4000")
}
