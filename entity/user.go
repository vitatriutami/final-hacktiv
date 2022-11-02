package entity

import (
	t "time"
)

type User struct {
	UserID      int `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt   t.Time
	UpdatedAt   t.Time
	Username    string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Password    string
	Age         int
	Photo       []Photo     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE" validate:"-"`
	SocialMedia SocialMedia `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE" validate:"-"`
	Comment     []Comment   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE" validate:"-"`
}
