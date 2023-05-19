package models

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoURL string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~URL of your photo is required"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
