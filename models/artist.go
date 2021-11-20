package models

import "time"

type Artist struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	Name       string     `json:name`
	Permalink  string     `json:permalink gorm:"uniqueIndex"`
	Bio        string     `json:bio`
	FormedAt   *time.Time `json:formed_at`
	Verified   bool       `json:verified`
	VerifiedAt time.Time  `json:verified_at`
	AvatarURL  string     `json:avatar_url`
	BannerURL  string     `json:banner_url`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at`
	UpdatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at`
}
