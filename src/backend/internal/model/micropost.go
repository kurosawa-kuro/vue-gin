package model

import (
	"time"
)

type Micropost struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null;size:200"`
	UserID    uint      `json:"userId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
