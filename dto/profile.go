package dto

import "waysbeans/models"

type ProfileResponse struct {
	ID      int                  `json:"id" gorm:"primary_key:auto_increment"`
	Phone   string               `json:"phone"`
	Address string               `json:"address"`
	UserID  int                  `json:"user_id"`
	User    models.UsersRelation `json:"user"`
}

type CreateProfileRequest struct {
	Phone   string `json:"phone" form:"phone" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
}
