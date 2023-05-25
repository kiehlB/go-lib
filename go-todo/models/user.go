package models

import "time"

type User struct {
	ID        int    `gorm:"primary_key:auto_increment"`
	Email     string `gorm:"type: varchar(255)"`
	Name      string `gorm:"type: varchar(255)"`
	Password  string `gorm:"type: varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Profile   ProfileResponse
}

type UsersProfileResponse struct {
	ID   int
	Name string
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
