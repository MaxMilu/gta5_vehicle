package model

type BaseModel struct {
	Like       bool `gorm:"size:1;not null;default:'0'"`
	Wishlist   bool `gorm:"size:1;not null;default:'0'"`
	AlreadyHas bool `gorm:"size:1;not null;default:'0'"`
}
