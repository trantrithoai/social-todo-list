package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social-todo-list/common"
	"social-todo-list/middleware"
	gin_item "social-todo-list/modules/item/transport/gin"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.Use(middleware.Recovery())

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", gin_item.CreateItem(db))
			items.GET("", gin_item.ListItem(db))
			items.GET("/:id", gin_item.GetItem(db))
			items.PATCH("/:id", gin_item.UpdateItem(db))
			items.DELETE("/:id", gin_item.DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000")
}
