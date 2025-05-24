package middleware

import (
	"go-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware is a middleware that checks if the user has the required role
type roles struct {
	RequiredRole string `json:"requiredRole"`
	UserRole     string `json:"userRole"`
}

var roleRank = map[string]int{
    "user":    1,
    "manager": 2,
    "admin":   3,
}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		// Debug log to see if the role is present
		if !exists {
			utils.Response(c, http.StatusForbidden, false, "Forbidden: Missing role", nil)
			c.Abort()
			return
		}

		// Type assertion to make sure the role is a string
		userRole, ok := role.(string)
		if !ok {
			utils.Response(c, http.StatusForbidden, false, "Forbidden: Invalid role type", nil)
			c.Abort()
			return
		}
		
		// ACL Logic
		userRank := roleRank[userRole]
		requiredRank := roleRank[requiredRole]

		// build response
		roles := roles{
			RequiredRole: requiredRole,
			UserRole:     userRole,
		}
		if userRank < requiredRank {
			utils.Response(c, http.StatusForbidden, false, "Forbidden: Insufficient permissions", roles)
			c.Abort()
			return
		}

		c.Next()
	}
}
