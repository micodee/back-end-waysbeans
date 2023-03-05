package models

import "time"

type User struct {
	ID        int                   `json:"id"`
	Name      string                `json:"fullname" gorm:"type: varchar(255)"`
	Email     string                `json:"email" gorm:"type: varchar(255)"`
	Password  string                `json:"password" gorm:"type: varchar(255)"`
	Profile   ProfileRelation       `json:"profile"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
}

type UsersToProfile struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UsersToProduct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UsersToCart struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersToProfile) TableName() string {
	return "users"
}

func (UsersToProduct) TableName() string {
	return "users"
}

func (UsersToCart) TableName() string {
	return "users"
}
