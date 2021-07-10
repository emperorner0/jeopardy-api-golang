package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	ShowNumber int    `json:"show_number"`
	Round      string `json:"round"`
	Category   string `json:"category"`
	Value      int    `json:"value"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
}
