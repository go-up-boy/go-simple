package routes

import (
	"go-simple/app/http/controllers/api/v1/user"
	"go-simple/globals"
)

func init() {
	r := globals.GlobalService.R
	group := r.Group("/api/v1/user")
	{
		var ctler = new(user.UsersController)
		group.GET("/", ctler.Index)
		group.GET("/:id", ctler.Show)
		group.POST("/", ctler.Store)
		group.PUT("/", ctler.Update)
		group.DELETE("/:id", ctler.Delete)
	}
}
