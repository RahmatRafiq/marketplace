package controllers

import (
	"errors"

	"golang_starter_kit_2025/app/helpers"
	"golang_starter_kit_2025/app/models" // Import models
	"golang_starter_kit_2025/app/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewPermissionController() *PermissionController {
	return &PermissionController{
		service: services.NewPermissionService(),
	}
}

type PermissionController struct {
	service *services.PermissionService
}

// @Summary		Get All Permissions
// @Description	API untuk mendapatkan semua Permission
// @Tags			Permission
// @Accept			json
// @Produce		json
// @Success		200	{object}	helpers.ResponseParams[models.Permission]{data=[]models.Permission}
// @Failure		500	{object}	map[string]string	"Internal Server Error"
// @Router			/permissions [get]
func (c *PermissionController) List(ctx *gin.Context) {
	permissions, err := c.service.GetAll()
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": err.Error()},
			Message:   "Gagal mendapatkan daftar Permission",
			Reference: "ERROR-3",
		}, 500)
		return
	}

	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.Permission]{Data: &permissions}, 200)
}

// @Summary		Create/Update Permission
// @Description	API untuk mengupdate atau membuat Permission
// @Tags			Permission
// @Accept			json
// @Produce		json
// @Param			permission	body		models.Permission	true	"Permission Data"
// @Success		200	{object}	helpers.ResponseParams[models.Permission]{item=models.Permission}
// @Failure		400	{object}	map[string]string	"Invalid input data"
// @Failure		500	{object}	map[string]string	"Internal Server Error"
// @Router			/permissions [put]
func (c *PermissionController) Put(ctx *gin.Context) {
	var permission models.Permission
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
				Errors:    helpers.ValidationError(verr),
				Message:   "Parameter tidak valid",
				Reference: "ERROR-4",
			}, 400)
			return
		}
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": err.Error()},
			Message:   "Gagal membuat Permission",
			Reference: "ERROR-3",
		}, 400)
		return
	}

	updatedPermission, err := c.service.Put(permission)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": err.Error()},
			Message:   "Gagal membuat Permission",
			Reference: "ERROR-3",
		}, 400)
		return
	}

	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.Permission]{Item: &updatedPermission}, 200)
}

// @Summary		Create/Update Permission
// @Description	API untuk mengupdate atau membuat Permission
// @Tags			Permission
// @Accept			json
// @Produce		json
// @Param			permission	body		models.Permission	true	"Permission Data"
// @Success		200	{object}	helpers.ResponseParams[models.Permission]{item=models.Permission}
// @Failure		400	{object}	map[string]string	"Invalid input data"
// @Failure		500	{object}	map[string]string	"Internal Server Error"
// @Router			/permissions [put]
func (c *PermissionController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Errors:    map[string]string{"error": err.Error()},
			Message:   "Gagal menghapus Permission",
			Reference: "ERROR-3",
		}, 400)
		return
	}

	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.Permission]{Message: "Permission deleted"}, 200)
}
