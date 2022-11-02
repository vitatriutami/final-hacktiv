package router

import (
	c "hacktiv-final-project/controller"
	g "github.com/gin-gonic/gin"
)

func PhotosRoutes(r *g.Engine, mid ...g.HandlerFunc) {
	rg := r.Group("/photos")
	rg.Use(mid[0])
	rg.GET("", c.GetAllPhotos)
	rg.POST("", mid[1], c.AddPhoto)
	rg.PUT("/:photoId", mid[1], c.UpdatePhoto)
	rg.DELETE("/:photoId", c.DeletePhoto)
}
