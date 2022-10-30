package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mksdziag/gorm_grocery/grocery"
	"github.com/mksdziag/gorm_grocery/model"
)

func main() {
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	db.DB()

	router := gin.Default()

	router.GET("/groceries", grocery.GetGroceries)
	router.GET("/groceries/:id", grocery.GetGrocery)
	router.POST("/groceries", grocery.PostGrocery)
	router.PUT("/groceries/:id", grocery.UpdateGrocery)
	router.DELETE("/groceries/:id", grocery.DeleteGrocery)

	log.Fatal(router.Run(":10000"))
}
