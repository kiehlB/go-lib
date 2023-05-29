package models

import "time"

type Post struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Content   string    `json:"content" gorm:"type:text"`
	UserID    int       `json:"userId"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
