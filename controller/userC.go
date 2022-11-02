package controller

import (
	e "hacktiv-final-project/entity"
	g "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	b "golang.org/x/crypto/bcrypt"
)

func Register(ctx *g.Context) {
	u := ctx.MustGet("user").(e.User)
	byt, _ := b.GenerateFromPassword([]byte(u.Password), b.DefaultCost)
	u.Password = string(byt)
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&u).Error
	if err != nil {
		ctx.JSON(400, g.H{"message": "Failed to register!"})
	} else {
		ctx.JSON(201, g.H{
			"id": u.UserID, "email": u.Email, "username": u.Username, "age": u.Age,
		})
	}
}

func Login(ctx *g.Context) {
	token := ctx.MustGet("token")
	ctx.JSON(200, g.H{"token":token})
}

func UpdateUser(ctx *g.Context) {
	sub := ctx.MustGet("sub").(map[string]interface{})
	u2 := e.User{UserID: int(sub["user_id"].(float64))}
	db := ctx.MustGet("db").(*gorm.DB)
	db.Take(&u2)
	u := ctx.MustGet("user").(e.User)
	u2.Email = u.Email
	u2.Username = u.Username
	db.Save(u2)
	ctx.JSON(200, g.H{
		"id": u2.UserID, "email": u2.Email, "username": u2.Username,
		"age": u2.Age, "updated_at": u2.UpdatedAt,
	})
}

func DeleteUser(ctx *g.Context) {
	sub := ctx.MustGet("sub").(map[string]interface{})
	userId := int(sub["user_id"].(float64))
	u := e.User{UserID: userId, Email: sub["email"].(string), Username: sub["username"].(string)}
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Delete(&u).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message":"User doesn't exists"})
		return
	}
	ctx.JSON(200, g.H{"message":"Your account has been deleted successfully"})
}
