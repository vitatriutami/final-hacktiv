package middleware

import (
	g "github.com/gin-gonic/gin"
	s "strings"
	b "golang.org/x/crypto/bcrypt"
	e "hacktiv-final-project/entity"
	util "hacktiv-final-project/utility"
	"gorm.io/gorm"
	"log"
)

func Authentication() g.HandlerFunc {
	return func(ctx *g.Context) {
		u := e.User{}
		ctx.BindJSON(&u)
		db := ctx.MustGet("db").(*gorm.DB)
		u2 := e.User{}
		db.Take(&u2, &e.User{Email: u.Email})
		err := b.CompareHashAndPassword([]byte(u2.Password), []byte(u.Password))
		if u.Email != u2.Email && err != nil {
			ctx.AbortWithStatusJSON(200, g.H{"message":"Email or Password are invalid!"})
		} else if u.Email != u2.Email {
			ctx.AbortWithStatusJSON(200, g.H{"message":"Email doesn't exists!"})
		} else if err != nil {
			ctx.AbortWithStatusJSON(200, g.H{"message":"Password is invalid!"})
		} else {
			log.Println("Login successfully!")
			data := map[string]any{
				"user_id": u2.UserID,
				"email": u2.Email,
				"username": u2.Username,
			}
			token := util.GenerateToken(data)
			ctx.Set("token", token)
			ctx.Next()
		}
	}
}

func Authorization() g.HandlerFunc {
	return func(ctx *g.Context) {
		header := ctx.Request.Header["Authorization"]
		ts := s.Replace(header[0], "Bearer ", "", -1)
		sub, err := util.ParseToken(ts)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(401, g.H{"message":err.Error()})
		} else {
			ctx.Set("sub", sub)
			ctx.Next()
		}
	}
}
