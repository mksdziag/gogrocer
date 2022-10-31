package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mksdziag/gorm_grocery/grocery"
	"github.com/mksdziag/gorm_grocery/middleware"
	"github.com/mksdziag/gorm_grocery/model"
	"github.com/mksdziag/gorm_grocery/token"
	"github.com/mksdziag/gorm_grocery/user"
)

func main() {
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	db.DB()

	router := gin.Default()

	api := router.Group("/api")

	api.POST("/token", token.GenerateToken)
	api.POST("/user/register", user.RegisterUser)

	groceries := api.Group("/groceries")
	groceries.Use(middleware.Auth())

	groceries.GET("/", grocery.GetGroceries)
	groceries.GET("/:id", grocery.GetGrocery)
	groceries.POST("/", grocery.PostGrocery)
	groceries.PUT("/:id", grocery.UpdateGrocery)
	groceries.DELETE("/:id", grocery.DeleteGrocery)

	log.Fatal(router.Run(":10000"))
}
