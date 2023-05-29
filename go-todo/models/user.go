package models

import "time"

type User struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Name      string          `json:"name" gorm:"type: varchar(255)"`
	Password  string          `json:"password" gorm:"type: varchar(255)"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Profile   ProfileResponse `json:"profile"`
	Posts     []Post          `json:"posts"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
