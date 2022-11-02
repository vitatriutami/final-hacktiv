package middleware

import (
	v "github.com/go-playground/validator/v10"
	g "github.com/gin-gonic/gin"
	e "hacktiv-final-project/entity"
	"log"
)

func RegisterUserValidation() g.HandlerFunc {
	return func(ctx *g.Context) {
		u := e.User{}
		ctx.BindJSON(&u)
		validate := v.New()
		validate.RegisterStructValidationMapRules(map[string]string{
			"Username": "required", "Email": "required,email",
			"Password": "required,min=6", "Age": "required,gt=8",
		}, e.User{})
		if err := validate.Struct(u); err != nil {
			if _, ok := err.(*v.InvalidValidationError); ok {
				log.Println(err)
				return
			}
			ctx.AbortWithStatusJSON(400, g.H{"message": err.Error()})
		} else {
			ctx.Set("user", u)
			ctx.Next()
		}
	}
}

func UpdateUserValidation() g.HandlerFunc {
	return func(ctx *g.Context) {
		u := e.User{}
		ctx.BindJSON(&u)
		validate := v.New()
		validate.RegisterStructValidationMapRules(map[string]string{
			"Username": "required", "Email": "required,email",
		}, e.User{})
		if err := validate.Struct(u); err != nil {
			if _, ok := err.(*v.InvalidValidationError); ok {
				log.Println(err)
				return
			}
			ctx.AbortWithStatusJSON(400, g.H{"message": err.Error()})
		} else {
			ctx.Set("user", u)
			ctx.Next()
		}
	}
}
