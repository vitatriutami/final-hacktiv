package entity

import (
	t "time"
)

type SocialMedia struct {
	SocialMediaID  int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	CreatedAt      t.Time `json:"createdAt"`
	UpdatedAt      t.Time `json:"updatedAt"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int    `gorm:"not null" json:"UserId"`
}
