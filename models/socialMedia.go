package models

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaURL string `gorm:"not null" json:"socialmediaurl" form:"socialmedia_url" valid:"required~Your social media url is required"`
	UserID         uint   `gorm:"not null" json:"user_id"`
}
