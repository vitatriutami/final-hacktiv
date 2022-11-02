package router

import (
	g "github.com/gin-gonic/gin"
	c "hacktiv-final-project/controller"
)

func UsersRoutes(r *g.Engine, mid ...g.HandlerFunc) {
	rg := r.Group("/users")
	rg.POST("/register", mid[2], c.Register)
	rg.POST("/login", mid[0], c.Login)
	rg.PUT("", mid[1], mid[3], c.UpdateUser)
	rg.DELETE("", mid[1], c.DeleteUser)
}
