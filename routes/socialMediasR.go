package router

import (
	g "github.com/gin-gonic/gin"
	c "hacktiv-final-project/controller"
)

func SocialMediasRoutes(r *g.Engine, mid ...g.HandlerFunc) {
	rg := r.Group("/socialmedias")
	rg.Use(mid[0])
	rg.GET("", c.GetAllSocialMedias)
	rg.POST("", mid[1], c.AddSocialMedia)
	rg.PUT("/:socialMediaId", mid[1], c.UpdateSocialMedia)
	rg.DELETE("/:socialMediaId", c.DeleteSocialMedia)
}
