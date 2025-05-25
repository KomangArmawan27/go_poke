package handlers

import (
	"go-api/config"
	"go-api/internal/dto"
	"go-api/internal/models"
	"go-api/internal/utils"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Get all pokemons

// Pokemon routes
// GetPokemons godoc
// @Summary      Get all pokemons
// @Description  Get list of pokemons with filters
// @Tags         Pokemons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        name     query     string  false  "Filter by name"
// @Param        type     query     string  false  "Filter by type"
// @Param        notes    query     string  false  "Filter by notes"
// @Param        sort_by  query     string  false  "Sort by field (name, type, notes)"
// @Param        order    query     string  false  "Sort order (asc or desc), default is asc"
// @Param        page     query     int     false  "Page number for pagination"
// @Param        limit    query     int     false  "Number of items per page"
// @Success      200    {object}  utils.BaseResponse
// @Failure      401    {object}  utils.BaseResponse  "Missing or invalid token"
// @Failure      403    {object}  utils.BaseResponse  "Forbidden access"
// @Failure      404    {object}  utils.BaseResponse  "Pokemons not found"
// @Router 		 /pokemons [get]
func GetPokemons(c *gin.Context) {
	// define the table
	var pokemons []models.Pokemon

	// fetching data from database
	db := config.DB

	// hanlde filter
	allowedField := map[string]string{
		"name":  "string",
		"type": "string",
		"notes":  "string",
	}
	db = utils.ApplyFilters(c, db, allowedField)

	// handle sort
	allowedSortFields := map[string]bool{
		"name":  true,
		"type":  true,
		"notes": true,
	}
	sortBy := c.Query("sort_by")
	order := c.DefaultQuery("order", "asc")
	if allowedSortFields[sortBy] {
		if order != "asc" && order != "desc" {
			order = "asc"
		}
		db = db.Order(fmt.Sprintf("%s %s", sortBy, order))
	}

	// get pagination
	db, pagination := utils.ApplyPagination(c, db, &models.Pokemon{})

	// fetch pokemon data
	if err := db.Find(&pokemons).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Failed to fetch pokemons", nil)
		return
	}

	if len(pokemons) == 0 {
		utils.Response(c, http.StatusNotFound, false, "Pokemons not found", nil)
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
		Items:           pokemons,
	}
	utils.Response(c, http.StatusOK, true, "Succes fetching pokemons data", dataResponse)
}

// Get Pokemon by ID
// GetPokemon godoc
// @Summary      Get pokemon by id
// @Description  Get list of pokemons by ID
// @Tags         Pokemons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   query     integer  false  "id"
// @Success      200    {object}  utils.BaseResponse
// @Failure      401    {object}  utils.BaseResponse  "Missing or invalid token"
// @Failure      403    {object}  utils.BaseResponse  "Forbidden access"
// @Failure      404    {object}  utils.BaseResponse  "Pokemon not found"
// @Router 		 /pokemon [get]
func GetPokemonByID(c *gin.Context) {
	id := c.Query("id")
	var pokemon models.Pokemon

	// error handling
	if err := config.DB.First(&pokemon, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Pokemon not found", nil)
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
		Items:           pokemon,
	}
	utils.Response(c, http.StatusOK, true, "Succes fetching pokemon data", dataResponse)
}

// Create Pokemon
// CreatePokemon godoc
// @Summary      Create new pokemon
// @Description  Creating new pokemon with role
// @Tags         Pokemons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        credentials body dto.CreateFavoritePokemonRequest true "Create pokemon credentials"
// @Success      200 {object} dto.Token
// @Failure      400 {object} utils.BaseResponse "Validation error"
// @Failure      401 {object} utils.BaseResponse "Missing or invalid token"
// @Failure      403 {object} utils.BaseResponse "Forbidden access"
// @Failure      409 {object} utils.BaseResponse "Pokemon email already used"
// @Failure      500 {object} utils.BaseResponse "Failed to hash password"
// @Router       /pokemon/create [post]
func CreatePokemon(c *gin.Context) {
	// get the serializer and validate it
	var req dto.CreateFavoritePokemonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// create the pokemon
	pokemon := models.Pokemon{
		Name:  req.Name,
		Type: req.Type,
		Notes:  req.Notes,
		Sprite:  req.Sprite,
	}

	// save the pokemon
	if err := config.DB.Create(&pokemon).Error; err != nil {
		utils.Response(c, http.StatusConflict, false, "Pokemon email already used", nil)
		return
	}

	utils.Response(c, http.StatusCreated, true, "Pokemon created", pokemon)
}

// Update Pokemon
// UpdatePokemon godoc
// @Summary      Update pokemon
// @Description  Update pokemon by id
// @Tags         Pokemons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   query     integer  false  "id"
// @Param        credentials body dto.UpdateFavoritePokemonRequest true "Update pokemon credentials"
// @Success      200 {object} dto.Token
// @Failure      401 {object}  utils.BaseResponse  "Missing or invalid token"
// @Failure      403 {object}  utils.BaseResponse  "Forbidden access"
// @Failure      404 {object}  utils.BaseResponse  "Pokemons not found"
// @Router       /pokemon/update [put]
func UpdatePokemon(c *gin.Context) {
	id := c.Query("id")

	// get the serializer and validate it
	var req dto.UpdateFavoritePokemonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// fetching pokemon data
	var pokemon = models.Pokemon{}
	if err := config.DB.First(&pokemon, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Pokemon not found", nil)
		return
	}

	// update the pokemon data
	pokemon.Name = req.Name
	pokemon.Type = req.Type
	pokemon.Notes = req.Notes

	// save the pokemon
	config.DB.Save(&pokemon)
	utils.Response(c, http.StatusOK, true, "Pokemon updated", pokemon)
}

// Delete Pokemon
// DeletePokemon godoc
// @Summary      Delete pokemon
// @Description  Delete pokemon by id
// @Tags         Pokemons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   query     integer  false  "id"
// @Success      200 {object} dto.Token
// @Failure      401 {object}  utils.BaseResponse  "Missing or invalid token"
// @Failure      403 {object}  utils.BaseResponse  "Forbidden access"
// @Failure      404 {object}  utils.BaseResponse  "Pokemons not found"
// @Router       /pokemon/delete [delete]
func DeletePokemon(c *gin.Context) {
	id := c.Query("id")
	var pokemon models.Pokemon

	if err := config.DB.First(&pokemon, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Pokemon not found", nil)
		return
	}

	config.DB.Delete(&pokemon)
	utils.Response(c, http.StatusOK, true, "Pokemon deleted", pokemon)
}
