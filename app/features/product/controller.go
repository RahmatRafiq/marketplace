package product

import (
	"errors"
	"golang_starter_kit_2025/app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	service *ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		service: NewProductService(),
	}
}

// @Summary		Get all products
// @Description	API untuk mendapatkan semua produk
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			request	query		FilterRequest	false	"Filter request"
// @Success		200		{object}	helpers.ResponseParams[Product]{data=[]Product}
// @Router			/products [get]
func (c *ProductController) GetAll(ctx *gin.Context) {
	var filters FilterRequest // Ganti dengan FilterRequest lokal jika sudah dimodularisasi
	if err := ctx.ShouldBindQuery(&filters); err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Parameter tidak valid",
			Reference: "ERROR-4",
		}, http.StatusBadRequest)
		return
	}

	products, err := c.service.GetAll(filters)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Gagal mendapatkan daftar produk",
			Reference: "ERROR-3",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusInternalServerError)
		return
	}

	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[Product]{Data: &products}, http.StatusOK)
}

// @Summary		Get product by ID
// @Description	API untuk mendapatkan produk berdasarkan ID
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Product ID"
// @Success		200	{object}	helpers.ResponseParams[Product]{item=Product}
// @Router			/products/{id} [get]
func (c *ProductController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := c.service.GetByID(id)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Gagal mendapatkan produk",
			Reference: "ERROR-2",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusNotFound)
		return
	}

	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[Product]{Item: &product}, http.StatusOK)
}

// @Summary		Create/Update product
// @Description	API untuk membuat atau mengupdate produk
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			product	body		ProductRequest	true	"Product request body"
// @Success		200		{object}	helpers.ResponseParams[Product]{item=Product}
// @Router			/products [put]
func (c *ProductController) Put(ctx *gin.Context) {
	var request ProductRequest
	if err := ctx.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
				Message:   "Periksa kembali form anda",
				Errors:    helpers.ValidationError(verr),
				Reference: "ERROR-4",
			}, http.StatusBadRequest)
			return
		}
	}

	product, err := c.service.Put(ctx, request)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Gagal membuat atau mengupdate produk",
			Reference: "ERROR-3",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusInternalServerError)
		return
	}

	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[Product]{Item: product}, http.StatusOK)
}

// @Summary		Delete product
// @Description	API untuk menghapus produk
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Product ID"
// @Success		200	{object}	helpers.ResponseParams[Product]
// @Router			/products/{id} [delete]
func (c *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Gagal menghapus produk",
			Reference: "ERROR-3",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusNotFound)
		return
	}

	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[any]{}, http.StatusOK)
}
