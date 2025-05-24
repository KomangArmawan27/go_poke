package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fmt"
	"strconv"
)

// ApplyFilters applies dynamic filters based on allowed fields
func ApplyFilters(c *gin.Context, db *gorm.DB, allowedFields map[string]string) *gorm.DB {
	for field, fieldType := range allowedFields {
		value := c.Query(field)
		if value == "" {
			continue
		}

		switch fieldType {
		case "string":
			db = db.Where(fmt.Sprintf("%s ILIKE ?", field), "%"+value+"%")

		case "int":
			if intValue, err := strconv.Atoi(value); err == nil {
				db = db.Where(fmt.Sprintf("%s = ?", field), intValue)
			}

		case "bool":
			if boolValue, err := strconv.ParseBool(value); err == nil {
				db = db.Where(fmt.Sprintf("%s = ?", field), boolValue)
			}

		case "date_from":
			db = db.Where(fmt.Sprintf("%s >= ?", field), value)

		case "date_to":
			db = db.Where(fmt.Sprintf("%s <= ?", field), value)
		}
	}
	return db
}

