package controller

import (
	e "hacktiv-final-project/entity"
	g "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	s "strconv"
)

func GetAllComments(ctx *g.Context) {
	co := []e.Comment{}
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Find(&co).Error
	if err != nil {
		ctx.AbortWithStatusJSON(200, err)
	} else {
		csms := make([]interface{}, len(co))
		for i, v := range co {
			csm := map[string]interface{}{}
			csm["id"] = v.CommentID
			csm["message"] = v.Message
			csm["photo_id"] = v.PhotoID
			csm["user_id"] = v.UserID
			csm["updated_at"] = v.UpdatedAt
			csm["created_at"] = v.CreatedAt
			u := e.User{UserID: v.UserID}
			db.First(&u)
			csm["User"] = map[string]interface{}{
				"id": u.UserID, "email": u.Email, "username": u.Username,
			}
			p := e.Photo{PhotoID: v.PhotoID}
			db.First(&p)
			csm["Photo"] = map[string]interface{}{
				"id": p.PhotoID,
				"title": p.Title,
				"caption": p.Caption,
				"photo_url": p.PhotoUrl,
				"user_id": p.UserID,
			}
			csms[i] = csm
		}
		ctx.JSON(200, csms)
	}
}

func AddComment(ctx *g.Context) {
	co := ctx.MustGet("comment").(e.Comment)
	sub := ctx.MustGet("sub").(map[string]interface{})
	co.UserID = int(sub["user_id"].(float64))
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&co).Error
	if err != nil {
		ctx.JSON(400, g.H{"message": "Failed to add new comment!"})
	} else {
		ctx.JSON(201, g.H{
			"id": co.CommentID, "message": co.Message, "photo_id": co.PhotoID,
			"user_id": co.UserID, "created_at": co.CreatedAt,
		})
	}
}

func UpdateComment(ctx *g.Context) {
	co := ctx.MustGet("comment").(e.Comment)
	cId, _ := s.Atoi(ctx.Param("commentId"))
	co.CommentID = cId
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Model(&co).Where("id = ?", co.CommentID).Updates(&co).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message": err.Error()})
	} else {
		db.Take(&co)
		p := e.Photo{PhotoID: co.PhotoID}
		db.Take(&p)
		ctx.JSON(200, g.H{
			"id": p.PhotoID, "title": p.Title, "caption": p.Caption,
			"photo_url": p.PhotoUrl, "user_id": p.UserID, "updated_at": co.UpdatedAt,
		})
	}
}

func DeleteComment(ctx *g.Context) {
	cId, _ := s.Atoi(ctx.Param("commentId"))
	co := e.Comment{CommentID: cId}
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Delete(&co).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message": "Comment doesn't exists"})
		return
	}
	ctx.JSON(200, g.H{"message": "Your comment has been deleted successfully"})
}
