package handlers

import (
	"go-api/config"
	"go-api/internal/models"
	"go-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts retrieves all products
func GetProducts(c *gin.Context) {
	// define the table
	var products []models.Product

	// fetching data from database
	db := config.DB

	// hanlde filter
	allowedField := map[string]string{
		"name":  "string",
		"price": "int",
	}
	db = utils.ApplyFilters(c, db, allowedField)

	// get pagination
	db, pagination := utils.ApplyPagination(c, db, &models.Product{})

	// fetch user data
	if err := db.Find(&products).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Failed to fetch products", nil)
		return
	}

	if len(products) == 0 {
		utils.Response(c, http.StatusNotFound, false, "Products not found", nil)
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
		Items:           products,
	}
	utils.Response(c, http.StatusOK, true, "Succes fetching products data", dataResponse)
}

// GetProductByID retrieves a single product by ID
func GetProductByID(c *gin.Context) {
	id := c.Query("id")
	var product models.Product

	// error handling
	if err := config.DB.First(&product, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Product not found", nil)
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
		Items:           product,
	}
	utils.Response(c, http.StatusOK, true, "Succes fetching product data", dataResponse)
}

// CreateProduct adds a new product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.Response(c, http.StatusBadRequest, false, "Wrong JSON format", nil)
		return
	}

	config.DB.Create(&product)
	utils.Response(c, http.StatusCreated, true, "Product created", product)
}

// UpdateProduct updates an existing product
func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Query("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Product not found", nil)
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.Response(c, http.StatusBadRequest, false, "Wrong JSON format", nil)
		return
	}

	config.DB.Save(&product)
	utils.Response(c, http.StatusOK, true, "Product updated", product)
}

// DeleteProduct removes a product by ID
func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Query("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "Product not found", nil)
		return
	}

	config.DB.Delete(&product)
	utils.Response(c, http.StatusOK, true, "Product deleted successfully", product)
}
