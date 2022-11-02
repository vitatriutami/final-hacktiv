package controller

import (
	e "hacktiv-final-project/entity"
	g "github.com/gin-gonic/gin"
	s "strconv"
	"gorm.io/gorm"
)

func GetAllSocialMedias(ctx *g.Context) {
	sm := []e.SocialMedia{}
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Find(&sm).Error
	if err != nil {
		ctx.AbortWithStatusJSON(200, err)
	} else {
		smm := make(map[string]interface{}, len(sm))
		for _, v := range sm {
			u := e.User{UserID: v.UserID}
			db.Take(&u)
			p := e.Photo{UserID: u.UserID}
			db.Take(&p)
			um := map[string]interface{}{
				"id": u.UserID, "username": u.Username,
				"profile_image_url": p.PhotoUrl,
			}
			smm["id"] = v.UserID
			smm["name"] = v.Name
			smm["social_media_url"] = v.SocialMediaUrl
			smm["User_id"] = v.UserID
			smm["createdAt"] = v.CreatedAt
			smm["updatedAt"] = v.UpdatedAt
			smm["User"] = um
		}
		ctx.JSON(200, g.H{"social_medias": smm})
	}
}

func AddSocialMedia(ctx *g.Context) {
	sm := ctx.MustGet("social_media").(e.SocialMedia)
	sub := ctx.MustGet("sub").(map[string]interface{})
	sm.UserID = int(sub["user_id"].(float64))
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&sm).Error
	if err != nil {
		ctx.JSON(400, g.H{"message": "Failed to add social media!"})
	} else {
		ctx.JSON(201, g.H{
			"id": sm.SocialMediaID, "name": sm.Name, "social_media_url": sm.SocialMediaUrl,
			"user_id": sm.UserID, "created_at": sm.CreatedAt,
		})
	}
}

func UpdateSocialMedia(ctx *g.Context) {
	sm := ctx.MustGet("social_media").(e.SocialMedia)
	smId, _ := s.Atoi(ctx.Param("socialMediaId"))
	sm.SocialMediaID = smId
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Model(&sm).Where("id = ?", sm.SocialMediaID).Updates(&sm).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message": err.Error()})
	} else {
		db.Take(&sm)
		ctx.JSON(200, g.H{
			"id": sm.SocialMediaID, "name": sm.Name, "social_media_url": sm.SocialMediaUrl,
			"user_id": sm.UserID, "updated_at": sm.UpdatedAt,
		})
	}
}

func DeleteSocialMedia(ctx *g.Context) {
	smId, _ := s.Atoi(ctx.Param("socialMediaId"))
	sm := e.SocialMedia{SocialMediaID: smId}
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Delete(&sm).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message": "Social media doesn't exists"})
		return
	}
	ctx.JSON(200, g.H{"message": "Your social media has been deleted successfully"})
}
