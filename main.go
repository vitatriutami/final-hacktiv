package main

import (
	conf "hacktiv-final-project/config"
	m "hacktiv-final-project/middleware"
	r "hacktiv-final-project/routes"
	g "github.com/gin-gonic/gin"
)

func main() {
	// g.SetMode(g.ReleaseMode)
	router := g.Default()
	router.Use(func(ctx *g.Context) {
		ctx.Set("db", conf.DB)
	})
	r.UsersRoutes(router, m.Authentication(), m.Authorization(), m.RegisterUserValidation(), m.UpdateUserValidation())
	r.PhotosRoutes(router, m.Authorization(), m.PhotoValidation())
	r.CommentsRoutes(router, m.Authorization(), m.CommentValidation(), m.UpdateCommentValidation())
	r.SocialMediasRoutes(router, m.Authorization(), m.SocialMediaValidation())
	router.Run()
}
