package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DataResponse struct {
	CurrentPage     int         `json:"currentPage"`
	TotalPages      int         `json:"totalPages"`
	TotalItems      int64       `json:"totalItems"`
	Limit           int         `json:"limit"`
	HasNextPage     bool        `json:"hasNextPage"`
	HasPreviousPage bool        `json:"hasPreviousPage"`
	Items           interface{} `json:"items"`
}

type Pagination struct {
	Limit           int
	Page            int
	Sort            string
	Offset          int
	HasNextPage     bool
	HasPreviousPage bool
	TotalItems      int64
	TotalPages      int
}

func ApplyPagination(c *gin.Context, db *gorm.DB, model interface{}) (*gorm.DB, Pagination) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	sort := c.DefaultQuery("sort", "id desc")

	var totalItems int64
	db.Model(model).Count(&totalItems)
	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))
	if totalPages == 0 {
		totalPages = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * limit

	pagination := Pagination{
		Limit:           limit,
		Page:            page,
		Sort:            sort,
		Offset:          offset,
		HasNextPage:     page < totalPages,
		HasPreviousPage: page > 1,
		TotalItems:      totalItems,
		TotalPages:      totalPages,
	}

	db = db.Order(sort).Limit(limit).Offset(offset)
	return db, pagination
}
