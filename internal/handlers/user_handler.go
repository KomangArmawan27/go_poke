package handlers

import (
	"go-api/config"
	"go-api/internal/dto"
	"go-api/internal/models"
	"go-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users

// User routes
// GetUsers godoc
// @Summary      Get all users
// @Description  Get list of users with filters
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        name   query     string  false  "Name filter"
// @Param        email  query     string  false  "Email filter"
// @Param        role   query     string  false  "Role filter"
// @Param        page     query     int     false  "Page number for pagination"
// @Param        limit    query     int     false  "Number of items per page"
// @Success      200    {object}  utils.BaseResponse
// @Failure      401    {object}  utils.BaseResponse  "Missing or invalid token"
// @Failure      403    {object}  utils.BaseResponse  "Forbidden access"
// @Failure      404    {object}  utils.BaseResponse  "Users not found"
// @Router 		 /users [get]
func GetUsers(c *gin.Context) {
	// define the table
	var users []models.User

	// fetching data from database
	db := config.DB

	// hanlde filter
	allowedField := map[string]string{
		"name":  "string",
		"email": "string",
		"role":  "string",
	}
	db = utils.ApplyFilters(c, db, allowedField)

	// get pagination
	db, pagination := utils.ApplyPagination(c, db, &models.User{})

	// fetch user data
	if err := db.Find(&users).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Failed to fetch users", nil)
		return
	}

	if len(users) == 0 {
		utils.Response(c, http.StatusNotFound, false, "Users not found", nil)
		return
	}

	// build the reponse
	dataResponse := utils.DataResponse{
		CurrentPage:     pagination.Page,
		TotalPages:      pagination.TotalPages,
		TotalItems:      pagination.TotalItems,
		Limit:           pagination.Limit,
		HasNextPage:     pagination.HasNextPage,
		HasPreviousPage: pagination.HasPreviousPage,
		Items:           users,
	}
	utils.Response(c, http.StatusOK, true, "Succes fetching users data", dataResponse)
}

// Get User by ID
// GetUser godoc
// @Summary      Get user by id
// @Description  Get list of users by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   query     integer  false  "id"
// @Success      200    {object}  utils.BaseResponse
// @Failure      401    {object}  utils.BaseResponse  "Missing or invalid token"
// @Failure      403    {object}  utils.BaseResponse  "Forbidden access"
// @Failure      404    {object}  utils.BaseResponse  "User not found"
// @Router 		 /user [get]
func GetUserByID(c *gin.Context) {
	id := c.Query("id")
	var user models.User

	// error handling
	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	// build the reponse
	dataResponse := utils.DataResponse{
		CurrentPage:     1,
		TotalPages:      1,
		TotalItems:      1,
		Limit:           1,
		HasNextPage:     false,
		HasPreviousPage: false,
		Items:           user,
	}
	utils.Response(c, http.StatusOK, true, "Succes fetching user data", dataResponse)
}

// Create User
// CreateUser godoc
// @Summary      Create new user
// @Description  Creating new user with role
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        credentials body dto.CreateUserRequest true "Create user credentials"
// @Success      200 {object} dto.Token
// @Failure      400 {object} utils.BaseResponse "Validation error"
// @Failure      401 {object} utils.BaseResponse "Missing or invalid token"
// @Failure      403 {object} utils.BaseResponse "Forbidden access"
// @Failure      409 {object} utils.BaseResponse "User email already used"
// @Failure      500 {object} utils.BaseResponse "Failed to hash password"
// @Router       /user/create [post]
func CreateUser(c *gin.Context) {
	// get the serializer and validate it
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// create the user
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// Hash the password before saving
	if err := user.HashPassword(req.Password); err != nil {
		utils.Response(c, http.StatusInternalServerError, false, "Failed to hash password", nil)
		return
	}

	// save the user
	if err := config.DB.Create(&user).Error; err != nil {
		utils.Response(c, http.StatusConflict, false, "User email already used", nil)
		return
	}

	utils.Response(c, http.StatusCreated, true, "User created", user)
}

// Update User
// UpdateUser godoc
// @Summary      Update user
// @Description  Update user by id
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   query     integer  false  "id"
// @Param        credentials body dto.UpdateUserRequest true "Update user credentials"
// @Success      200 {object} dto.Token
// @Failure      401 {object}  utils.BaseResponse  "Missing or invalid token"
// @Failure      403 {object}  utils.BaseResponse  "Forbidden access"
// @Failure      404 {object}  utils.BaseResponse  "Users not found"
// @Router       /user/update [put]
func UpdateUser(c *gin.Context) {
	id := c.Query("id")

	// get the serializer and validate it
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// fetching user data
	var user = models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	// update the user data
	user.Name = req.Name
	user.Email = req.Email
	user.Role = req.Role

	// save the user
	config.DB.Save(&user)
	utils.Response(c, http.StatusOK, true, "User updated", user)
}

// Delete User
// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete user by id
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   query     integer  false  "id"
// @Success      200 {object} dto.Token
// @Failure      401 {object}  utils.BaseResponse  "Missing or invalid token"
// @Failure      403 {object}  utils.BaseResponse  "Forbidden access"
// @Failure      404 {object}  utils.BaseResponse  "Users not found"
// @Router       /user/delete [delete]
func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	config.DB.Delete(&user)
	utils.Response(c, http.StatusOK, true, "User deleted", user)
}
