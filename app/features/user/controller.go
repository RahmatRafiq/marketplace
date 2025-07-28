package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssignRolesRequest struct {
	Roles []uint `json:"roles"`
}

func NewUserController() *UserController {
	return &UserController{
		service: NewUserService(),
	}
}

type UserController struct {
	service *UserService
}

// @Summary      List all users
// @Description  Retrieve a list of all users
// @Tags         users
// @Produce      json
// @Success      200 {array} User "List of users"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users [get]
func (c *UserController) List(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// @Summary      Get a user by ID
// @Description  Retrieve a user by their ID
// @Tags         users
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} User "User object"
// @Failure      404 {object} map[string]string "User not found"
// @Router       /users/{id} [get]
func (c *UserController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.Find(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// @Summary      Create or update a user
// @Description  Create a new user or update an existing one by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body User true "User object"
// @Success      200 {object} User "Created or updated user"
// @Failure      400 {object} map[string]string "Invalid input data"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users [put]
func (c *UserController) Put(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := c.service.Put(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

// @Summary      Delete a user by ID
// @Description  Delete a specific user by their ID
// @Tags         users
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} map[string]string "User deleted successfully"
// @Failure      404 {object} map[string]string "User not found"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users/{id} [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// @Summary      Assign roles to a user
// @Description  Assign one or more roles to a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        roles body AssignRolesRequest true "Roles to assign"
// @Success      200 {object} map[string]string "Roles assigned"
// @Failure      400 {object} map[string]string "Invalid input data"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users/{id}/roles [post]
func (c *UserController) AssignRoles(ctx *gin.Context) {
	var req AssignRolesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.Param("id")
	err := c.service.AssignRolesToUser(userId, req.Roles)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Roles assigned to user"})
}

// @Summary      Get roles of a user
// @Description  Retrieve all roles assigned to a user
// @Tags         users
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {array} models.Role "List of roles"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /users/{id}/roles [get]
func (c *UserController) GetRoles(ctx *gin.Context) {
	userId := ctx.Param("id")
	roles, err := c.service.GetRolesByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, roles)
}
