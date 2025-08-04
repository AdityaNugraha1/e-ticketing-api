package models

import "time"

type Admin struct {
	AdminID      uint      `gorm:"primary_key" json:"admin_id"`
	Username     string    `gorm:"type:varchar(50);unique_index" json:"username"`
	PasswordHash string    `gorm:"type:varchar(255)" json:"-"`
	Role         string    `gorm:"type:varchar(20)" json:"role"`
	LastSync     time.Time `json:"last_sync"`
}

type Terminal struct {
	TerminalID uint   `gorm:"primary_key" json:"terminal_id"`
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
}