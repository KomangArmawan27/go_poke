package models

import "gorm.io/gorm"

// Pokemon represents a pokemon entity
type Pokemon struct {
	gorm.Model
	UserEmail string `json:"userEmail"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Notes     string `json:"notes"`
	Sprite    string `json:"sprite"`
}
