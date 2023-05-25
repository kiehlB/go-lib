package models

import "time"

type Profile struct {
	ID        int                  `gorm:"primary_key:auto_increment"`
	Bio 	  string			   `gorm:"type: string"`
	UserID    int                  `gorm:"type: int"`
	User      UsersProfileResponse  
	CreatedAt time.Time            
	UpdatedAt time.Time           
}

type ProfileResponse struct {
	Bio    string  
	UserID  int     
}

func (ProfileResponse) TableName() string {
	return "profiles"
}