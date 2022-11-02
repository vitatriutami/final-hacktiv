package middleware

import (
	v "github.com/go-playground/validator/v10"
	g "github.com/gin-gonic/gin"
	e "hacktiv-final-project/entity"
	"log"
)

func SocialMediaValidation() g.HandlerFunc {
	return func(ctx *g.Context) {
		sm := e.SocialMedia{}
		ctx.BindJSON(&sm)
		validate := v.New()
		validate.RegisterStructValidationMapRules(map[string]string{
			"Name": "required", "SocialMediaUrl": "required",
		}, e.SocialMedia{})
		if err := validate.Struct(sm); err != nil {
			if _, ok := err.(*v.InvalidValidationError); ok {
				log.Println(err)
				return
			}
			ctx.AbortWithStatusJSON(400, g.H{"message": err.Error()})
		} else {
			ctx.Set("social_media", sm)
			ctx.Next()
		}
	}
}
