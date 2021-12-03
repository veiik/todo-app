package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signUp", h.SignUp)
		auth.POST("/signIn", h.SignIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.POST("/", h.createList)
			lists.DELETE("/:id", h.deleteList)
			lists.PUT("/:id", h.updateList)

			items := lists.Group(":id/items")
			{
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.POST("/", h.createItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
