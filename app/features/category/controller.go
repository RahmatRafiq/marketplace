package category

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewCategoryController() *CategoryController {
	return &CategoryController{
		service: NewCategoryService(),
	}
}

type CategoryController struct {
	service *CategoryService
}

// @Summary      List all categories
// @Description  Retrieve a list of all categories, including related products
// @Tags         categories
// @Security     Bearer
// @Produce      json
// @Success      200 {array} Category "List of categories with products"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /categories [get]
func (c *CategoryController) List(ctx *gin.Context) {
	categories, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// @Summary      Get a category by ID
// @Description  Retrieve a category by its ID, including related products
// @Tags         categories
// @Security     Bearer
// @Produce      json
// @Param        id path string true "Category ID"
// @Success      200 {object} Category "Category with products"
// @Failure      404 {object} map[string]string "Category not found"
// @Router       /categories/{id} [get]
func (c *CategoryController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	category, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// @Summary      Create or update a category
// @Description  Create a new category or update an existing one by ID
// @Tags         categories
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        category body Category true "Category Data"
// @Success      200 {object} Category "Created or updated category"
// @Failure      400 {object} map[string]string "Invalid input data"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /categories [put]
func (c *CategoryController) Put(ctx *gin.Context) {
	var req Category
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UpdatedAt = time.Now()
	updatedCategory, err := c.service.Put(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedCategory)
}

// @Summary      Delete a category by ID
// @Description  Delete a specific category by its ID
// @Tags         categories
// @Security     Bearer
// @Produce      json
// @Param        id path string true "Category ID"
// @Success      200 {object} map[string]string "Category deleted successfully"
// @Failure      404 {object} map[string]string "Category not found"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /categories/{id} [delete]
func (c *CategoryController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
