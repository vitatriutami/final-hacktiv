package entity

import (
	t "time"
)

type Comment struct {
	CommentID int `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt t.Time
  UpdatedAt t.Time
  Message string `validate:"required" json:"message"`
	UserID int `gorm:"not null" json:"user_id"`
	PhotoID int `gorm:"not null" json:"photo_id"`
}
