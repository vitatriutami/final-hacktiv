package controller

import (
	e "hacktiv-final-project/entity"
	g "github.com/gin-gonic/gin"
	s "strconv"
	"gorm.io/gorm"
)

func GetAllPhotos(ctx *g.Context) {
	ps := []e.Photo{}
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Find(&ps).Error
	if err != nil {
		ctx.AbortWithStatusJSON(200, err)
	} else {
		psms := make([]interface{}, len(ps))
		for i, v := range ps {
			psm := map[string]interface{}{}
			psm["id"] = v.PhotoID
			psm["title"] = v.Title
			psm["caption"] = v.PhotoUrl
			psm["user_id"] = v.UserID
			psm["created_at"] = v.CreatedAt
			psm["updated_at"] = v.UpdatedAt
			u := e.User{UserID: v.UserID}
			db.First(&u)
			um := map[string]interface{}{
				"email": u.Email, "username": u.Username,
			}
			psm["User"] = um
			psms[i] = psm
		}
		ctx.JSON(200, psms)
	}
}

func AddPhoto(ctx *g.Context) {
	p := ctx.MustGet("photo").(e.Photo)
	sub := ctx.MustGet("sub").(map[string]interface{})
	p.UserID = int(sub["user_id"].(float64))
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&p).Error
	if err != nil {
		ctx.JSON(400, g.H{"message": "Failed to add new photo!"})
	} else {
		ctx.JSON(201, g.H{
			"id": p.UserID, "title": p.Title, "caption": p.Caption,
			"photo_url": p.PhotoUrl, "created_at": p.CreatedAt,
		})
	}
}

func UpdatePhoto(ctx *g.Context) {
	p := ctx.MustGet("photo").(e.Photo)
	pId, _ := s.Atoi(ctx.Param("photoId"))
	p.PhotoID = pId
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Model(&p).Where("id = ?", p.PhotoID).Updates(&p).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message": err.Error()})
	} else {
		db.Take(&p)
		ctx.JSON(200, g.H{
			"id": p.PhotoID, "title": p.Title, "caption": p.Caption,
			"photo_url": p.PhotoUrl, "user_id": p.UserID, "updated_at": p.UpdatedAt,
		})
	}
}

func DeletePhoto(ctx *g.Context) {
	pId, _ := s.Atoi(ctx.Param("photoId"))
	p := e.Photo{PhotoID: pId}
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Delete(&p).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message": "Photo doesn't exists"})
		return
	}
	ctx.JSON(200, g.H{"message": "Your photo has been deleted successfully"})
}
