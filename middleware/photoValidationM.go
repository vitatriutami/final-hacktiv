package middleware

import (
	e "hacktiv-final-project/entity"
	"log"
	g "github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

func PhotoValidation() g.HandlerFunc {
	return func(ctx *g.Context) {
		p := e.Photo{}
		ctx.BindJSON(&p)
		validate := v.New()
		validate.RegisterStructValidationMapRules(map[string]string{
			"Title": "required", "Caption": "required",
			"PhotoUrl": "required",
		}, e.Photo{})
		if err := validate.Struct(p); err != nil {
			if _, ok := err.(*v.InvalidValidationError); ok {
				log.Println(err)
				return
			}
			ctx.AbortWithStatusJSON(400, g.H{"message": err.Error()})
		} else {
			ctx.Set("photo", p)
			ctx.Next()
		}
	}
}
