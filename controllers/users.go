package controllers

import (
	"net/http"
	"notasGo/models"
	"notasGo/services"
	"notasGo/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// GetUsers godoc
// @Summary Obtiene todos los usuarios
// @Description Devuelve una lista de todos los usuarios registrados
// @Tags usuarios
// @Produce json
// @Success 200 {object} models.UsersListResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, total, err := ctrl.userService.GetAllUsers()
	if err != nil {
		utils.InternalServerError(c, "Error al obtener usuarios", err)
		return
	}

	// Convert to response format
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	response := models.UsersListResponse{
		Success: true,
		Message: "Usuarios obtenidos exitosamente",
		Users:   userResponses,
		Total:   total,
	}

	c.JSON(http.StatusOK, response)
}

// GetUserByID godoc
// @Summary Obtiene un usuario por ID
// @Description Devuelve la información de un usuario específico
// @Tags usuarios
// @Produce json
// @Param id path int true "ID del usuario"
// @Success 200 {object} models.APIResponse{data=models.UserResponse}
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [get]
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	
	user, err := ctrl.userService.GetUserByID(id)
	if err != nil {
		if err.Error() == "usuario no encontrado" {
			utils.NotFoundError(c, "Usuario no encontrado")
			return
		}
		utils.InternalServerError(c, "Error al obtener usuario", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuario obtenido exitosamente", userResponse)
}

// RegisterUser godoc
// @Summary Registra un nuevo usuario
// @Description Crea una nueva cuenta de usuario con encriptación de contraseña
// @Tags usuarios
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "Datos del usuario"
// @Success 201 {object} models.APIResponse{data=models.UserResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /register [post]
func (ctrl *UserController) RegisterUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestError(c, "Datos inválidos", err)
		return
	}

	user, err := ctrl.userService.CreateUser(&req)
	if err != nil {
		if err.Error() == "el email ya está registrado" || err.Error() == "el nombre de usuario ya está en uso" {
			utils.ConflictError(c, err.Error(), nil)
			return
		}
		utils.InternalServerError(c, "Error al crear usuario", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusCreated, "Usuario registrado exitosamente", userResponse)
}

// UpdateUser godoc
// @Summary Actualiza un usuario
// @Description Actualiza la información de un usuario existente
// @Tags usuarios
// @Accept json
// @Produce json
// @Param id path int true "ID del usuario"
// @Param user body models.UpdateUserRequest true "Datos a actualizar"
// @Success 200 {object} models.APIResponse{data=models.UserResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestError(c, "Datos inválidos", err)
		return
	}

	user, err := ctrl.userService.UpdateUser(id, &req)
	if err != nil {
		if err.Error() == "usuario no encontrado" {
			utils.NotFoundError(c, "Usuario no encontrado")
			return
		}
		if err.Error() == "el email ya está en uso por otro usuario" || err.Error() == "el nombre de usuario ya está en uso" {
			utils.ConflictError(c, err.Error(), nil)
			return
		}
		utils.InternalServerError(c, "Error al actualizar usuario", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuario actualizado exitosamente", userResponse)
}

// DeleteUser godoc
// @Summary Elimina un usuario
// @Description Elimina un usuario y todas sus notas asociadas
// @Tags usuarios
// @Produce json
// @Param id path int true "ID del usuario"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	
	err := ctrl.userService.DeleteUser(id)
	if err != nil {
		if err.Error() == "usuario no encontrado" {
			utils.NotFoundError(c, "Usuario no encontrado")
			return
		}
		utils.InternalServerError(c, "Error al eliminar usuario", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuario y sus notas eliminados exitosamente", nil)
}

// LoginUser godoc
// @Summary Autenticación de usuario
// @Description Autentica un usuario con email y contraseña
// @Tags usuarios
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Credenciales de acceso"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /login [post]
func (ctrl *UserController) LoginUser(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestError(c, "Datos inválidos", err)
		return
	}

	user, err := ctrl.userService.AuthenticateUser(&req)
	if err != nil {
		if err.Error() == "credenciales inválidas" || err.Error() == "cuenta inactiva" {
			utils.UnauthorizedError(c, err.Error())
			return
		}
		utils.InternalServerError(c, "Error en el proceso de autenticación", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response := models.LoginResponse{
		Success: true,
		Message: "Login exitoso",
		User:    userResponse,
	}

	c.JSON(http.StatusOK, response)
}