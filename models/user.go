package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname" gorm:"type: varchar(255)"`
	Username  string    `json:"user_name" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type : varchar(255)"`
	Gender    string    `json:"gender" gorm:"type : varchar(255)"`
	Phone     string    `json:"phone" gorm:"type : varchar(255)"`
	Address   string    `json:"address" gorm:"type : text"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	ListAsID  int       `json:"list_as_id" gorm:"foreignKey :ID"`
	ListAs    ListAs    `json:"list_as"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type UserProfileResponse struct {
	ID       int            `json:"id"`
	Fullname string         `json:"fullname"`
	Username string         `json:"user_name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Gender   string         `json:"gender"`
	Phone    string         `json:"phone"`
	Address  string         `json:"address"`
	Image    string         `json:"image"`
	ListAs   ListAsResponse `json:"list_as"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}
