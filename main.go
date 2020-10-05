package main

import (
	"github.com/gin-gonic/gin"
	apiController "github.com/irvanherz/goblog/controller/api"
	blogController "github.com/irvanherz/goblog/controller/blog"
	"github.com/irvanherz/goblog/database"
	"github.com/irvanherz/goblog/middleware"
	"github.com/irvanherz/goblog/service"
)

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "Article, OPTIONS, GET, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
func main() {
	router := gin.Default()
	router.Use(corsMiddleware)
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("view/*")
	connection := database.NewDB()
	ps := service.NewArticleService(connection)
	us := service.NewUserService(connection)
	bac := blogController.NewArticleController(ps)
	aac := apiController.NewArticleController(ps)
	auc := apiController.NewAuthController(us)

	// BLOG ROUTERS
	router.GET("/", bac.Get)
	router.GET("/articles/:ID", bac.GetID)

	//API ROUTERS
	router.GET("/api/v1/articles", aac.Get)
	router.GET("/api/v1/articles/:ID", aac.GetID)
	router.POST("/api/v1/articles", middleware.AuthMiddleware(), aac.Post)
	router.PUT("/api/v1/articles/:ID", middleware.AuthMiddleware(), aac.PutID)
	router.DELETE("/api/v1/articles/:ID", middleware.AuthMiddleware(), aac.DeleteID)

	router.POST("/api/v1/auth/login", auc.Login)

	router.Run(":8080")
}
