package middleware

import (
	e "hacktiv-final-project/entity"
	"log"
	g "github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

func CommentValidation() g.HandlerFunc {
	return func(ctx *g.Context) {
		co := e.Comment{}
		ctx.BindJSON(&co)
		validate := v.New()
		validate.RegisterStructValidationMapRules(map[string]string{
			"Message": "required", "PhotoID": "required",
		}, e.Comment{})
		if err := validate.Struct(co); err != nil {
			if _, ok := err.(*v.InvalidValidationError); ok {
				log.Println(err)
				return
			}
			ctx.AbortWithStatusJSON(400, g.H{"message": err.Error()})
		} else {
			ctx.Set("comment", co)
			ctx.Next()
		}
	}
}

func UpdateCommentValidation() g.HandlerFunc {
	return func(ctx *g.Context) {
		co := e.Comment{}
		ctx.BindJSON(&co)
		validate := v.New()
		validate.RegisterStructValidationMapRules(map[string]string{
			"Message": "required",
		}, e.Comment{})
		if err := validate.Struct(co); err != nil {
			if _, ok := err.(*v.InvalidValidationError); ok {
				log.Println(err)
				return
			}
			ctx.AbortWithStatusJSON(400, g.H{"message": err.Error()})
		} else {
			ctx.Set("comment", co)
			ctx.Next()
		}
	}
}
