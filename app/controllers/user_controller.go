package controllers

import (
	"golang_starter_kit_2025/app/helpers"
	"golang_starter_kit_2025/app/models"
	"golang_starter_kit_2025/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssignRolesRequest struct {
	Roles []uint `json:"roles"`
}

func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserService(),
	}
}

type UserController struct {
	service *services.UserService
}

// @Summary      List all users
// @Description  Retrieve a list of all users
// @Tags         users
// @Produce      json
// @Success      200 {array} models.User "List of users"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users [get]
func (c *UserController) List(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Gagal mengambil data user",
			Reference: "USER-1",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusInternalServerError)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.User]{Data: &users}, http.StatusOK)
}

// @Summary      Get a user by ID
// @Description  Retrieve a user by their ID
// @Tags         users
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} models.User "User object"
// @Failure      404 {object} map[string]string "User not found"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users/{id} [get]
func (c *UserController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.Find(id)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "User tidak ditemukan",
			Reference: "USER-2",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusNotFound)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.User]{Item: &user}, http.StatusOK)
}

// @Summary      Create or update a user
// @Description  Create a new user or update an existing one by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "User object"
// @Success      200 {object} models.User "Created or updated user"
// @Failure      400 {object} map[string]string "Invalid input data"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users [put]
func (c *UserController) Put(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Data tidak valid",
			Reference: "USER-3",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusBadRequest)
		return
	}
	updatedUser, err := c.service.Put(user)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Gagal menyimpan user",
			Reference: "USER-4",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusInternalServerError)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.User]{Item: &updatedUser}, http.StatusOK)
}

// @Summary      Delete a user by ID
// @Description  Delete a specific user by their ID
// @Tags         users
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} any "User deleted successfully"
// @Failure      404 {object} map[string]string "User not found"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users/{id} [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "User tidak ditemukan",
			Reference: "USER-5",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusNotFound)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[any]{}, http.StatusOK)
}

// @Summary      Assign roles to a user
// @Description  Assign one or more roles to a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        roles body AssignRolesRequest true "Roles to assign"
// @Success      200 {object} any "Roles assigned successfully"
// @Failure      400 {object} map[string]string "Invalid input data"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users/{id}/roles [post]
func (c *UserController) AssignRoles(ctx *gin.Context) {
	var req AssignRolesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Data tidak valid",
			Reference: "USER-6",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusBadRequest)
		return
	}

	userId := ctx.Param("id")
	err := c.service.AssignRolesToUser(userId, req.Roles)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Gagal assign roles",
			Reference: "USER-7",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusInternalServerError)
		return
	}

	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[any]{}, http.StatusOK)
}

// @Summary      Get roles of a user
// @Description  Retrieve all roles assigned to a user
// @Tags         users
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {array} models.Role "List of roles"
// @Failure      404 {object} map[string]string "User not found"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users/{id}/roles [get]
func (c *UserController) GetRoles(ctx *gin.Context) {
	userId := ctx.Param("id")
	roles, err := c.service.GetRolesByUserId(userId)
	if err != nil {
		helpers.ResponseError(ctx, &helpers.ResponseParams[any]{
			Message:   "Gagal mengambil roles user",
			Reference: "USER-8",
			Errors:    map[string]string{"error": err.Error()},
		}, http.StatusInternalServerError)
		return
	}
	helpers.ResponseSuccess(ctx, &helpers.ResponseParams[models.Role]{Data: &roles}, http.StatusOK)
}
