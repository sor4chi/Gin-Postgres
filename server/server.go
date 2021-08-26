package server

import (
	"github.com/gin-gonic/gin"
	user "github.com/kawamurasorachi/Gin-Postgres/controller"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	ctrl := user.Controller{}

	u := r.Group("/users")
	{
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
