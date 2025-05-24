package handlers

import (
	"go-api/config"
	"go-api/internal/auth"
	"go-api/internal/dto"
	"go-api/internal/models"
	"go-api/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoginHandler authenticates users and generates JWT
// Login godoc
// @Summary      User Login
// @Description  Authenticates user and returns JWT token
// @Tags         OAuth
// @Accept       json
// @Produce      json
// @Param        credentials body dto.LoginRequest true "Login credentials"
// @Success      200 {object} dto.Token
// @Failure      401 {object} utils.BaseResponse
// @Router       /login [post]
func LoginHandler(c *gin.Context) {
	loginRequest := dto.LoginRequest{}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, false, "Invalid request", nil)
		return
	}

	var user models.User
	// hashing the password
	if err := user.HashPassword(loginRequest.Password); err != nil {
		utils.Response(c, http.StatusInternalServerError, false, "Failed to hash password", nil)
		return
	}

	if err := config.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		utils.Response(c, http.StatusUnauthorized, false, "Invalid email", nil)
		return
	}

	if !user.CheckPassword(loginRequest.Password) {
		utils.Response(c, http.StatusUnauthorized, false, "Invalid password", nil)
		return
	}

	// Generate JWT
	token, err := auth.GenerateToken(loginRequest.Email, user.Role)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, false, "Failed to generate token", nil)
		return
	}

	// build the reponse
	dateTime := dto.TimeJSON{Time: time.Now().Add(1 * time.Hour)}
	dataResponse := dto.Token{
		Username:    user.Name,
		Email:       user.Email,
		ActiveUntil: dateTime,
		Token:       token,
	}
	utils.Response(c, http.StatusOK, true, "Success to generate token", dataResponse)
}

// Register
// Register godoc
// @Summary      Register new user
// @Description  Register new user for member
// @Tags         OAuth
// @Accept       json
// @Produce      json
// @Param        credentials body dto.RegisterRequest true "Create user credentials"
// @Success      200 {object} dto.Token
// @Failure      400 {object} utils.BaseResponse "Validation error"
// @Failure      409 {object} utils.BaseResponse "User email already used"
// @Failure      500 {object} utils.BaseResponse "Failed to hash password"
// @Router       /register [post]
func RegisterHandler(c *gin.Context) {
	// get the serializer and validate it
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// create the user
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  "user",
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
