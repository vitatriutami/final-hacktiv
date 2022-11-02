package router

import (
	g "github.com/gin-gonic/gin"
	c "hacktiv-final-project/controller"
)

func CommentsRoutes(r *g.Engine, mid ...g.HandlerFunc) {
	rg := r.Group("/comments")
	rg.Use(mid[0])
	rg.GET("", c.GetAllComments)
	rg.POST("", mid[1], c.AddComment)
	rg.PUT("/:commentId", mid[2], c.UpdateComment)
	rg.DELETE("/:commentId", c.DeleteComment)
}
