package dto

import (
	"time"
	"fmt"
)

// token struct
type Token struct {
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	ActiveUntil TimeJSON `json:"activeUntil"`
	Token       string   `json:"token"`
}

// Custom time type that formats JSON output
type TimeJSON struct {
	time.Time
}

// Custom JSON marshaling to format time
func (t TimeJSON) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}