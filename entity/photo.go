package entity

import (
	t "time"
)

type Photo struct {
	PhotoID   int `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt t.Time
	UpdatedAt t.Time
	Title     string
	Caption   string
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `gorm:"not null"`
	Comment   []Comment `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE" validate:"-"`
}
