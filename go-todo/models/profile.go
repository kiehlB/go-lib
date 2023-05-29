package models

import "time"

type Profile struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Bio       string               `json:"bio" gorm:"type: varchar(255)"`
	UserID    int                  `json:"userId" gorm:"type: int"`
	User      UsersProfileResponse `json:"user"`
	CreatedAt time.Time            `json:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt"`
}

type ProfileResponse struct {
	Bio    string `json:"bio"`
	UserID int    `json:"userId"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
