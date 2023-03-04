package dto

import "waysbeans/models"

type ProfileResponse struct {
	ID      int                         `json:"id" gorm:"primary_key:auto_increment"`
	Phone   string                      `json:"phone" gorm:"type: varchar(255)"`
	Gender  string                      `json:"gender" gorm:"type: varchar(255)"`
	Address string                      `json:"address" gorm:"type: text"`
	UserID  int                         `json:"user_id"`
	User    models.UsersProfileResponse `json:"user"`
}

type CreateProfileRequest struct {
	Phone   string `json:"phone" form:"phone" validate:"required"`
	Gender  string `json:"gender" form:"gender" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
}
