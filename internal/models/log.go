package models

import "time"

// models/request_log.go
type Log struct {
	ID           uint      `gorm:"primaryKey"`
	Username     *uint     `json:"username"`
	Method       string    `json:"method"`
	URI          string    `json:"uri"`
	ClientIP     string    `json:"client_ip"`
	StatusCode   int       `json:"status_code"`
	Duration     string    `json:"duration"`
	RequestBody  string    `json:"request_body"`
	ResponseBody string    `json:"response_body"`
	CreatedAt    time.Time `json:"created_at"`
}
