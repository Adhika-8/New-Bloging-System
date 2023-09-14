package main

import (
	db "new-bloging-system/app"
	"new-bloging-system/controller"
	"new-bloging-system/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()
	db.DBmigrate()
	r := gin.Default()
	r.Use(middleware.PanicHandling())
	v1 := r.Group("/blog")
	{
		v1.POST("/create", controller.CreateBlogHandler)
		v1.GET("/get/:id", controller.GetBlogByIdHandler)
		v1.GET("/get/all", controller.GetAllBlogHandler)
		v1.DELETE("/delete/:id", controller.DeleteBlogByIdHandler)
		v1.PATCH("/update/:id", controller.UpdateBlogByIdHandler)
	}
	r.Run("localhost:8082")

}
