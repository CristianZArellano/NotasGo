package services

import (
	"errors"
	"notasGo/database"
	"notasGo/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// GetAllUsers retrieves all users from database
func (s *UserService) GetAllUsers() ([]models.User, int64, error) {
	var users []models.User
	var count int64
	
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, 0, err
	}
	
	database.DB.Model(&models.User{}).Count(&count)
	return users, count, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user with hashed password
func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// Check if email already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("el email ya está registrado")
	}
	
	// Check if username already exists
	if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("el nombre de usuario ya está en uso")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error al encriptar contraseña")
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
		Status:   "activo",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	// Clear password from response
	user.Password = ""
	return &user, nil
}

// UpdateUser updates user information
func (s *UserService) UpdateUser(id string, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Check for email conflicts if email is being updated
	if req.Email != "" && req.Email != user.Email {
		var existingUser models.User
		if err := database.DB.Where("email = ? AND id != ?", req.Email, id).First(&existingUser).Error; err == nil {
			return nil, errors.New("el email ya está en uso por otro usuario")
		}
	}

	// Check for username conflicts if username is being updated
	if req.Username != "" && req.Username != user.Username {
		var existingUser models.User
		if err := database.DB.Where("username = ? AND id != ?", req.Username, id).First(&existingUser).Error; err == nil {
			return nil, errors.New("el nombre de usuario ya está en uso")
		}
	}

	// Update fields
	updates := make(map[string]interface{})
	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Role != "" {
		updates["role"] = req.Role
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	if err := database.DB.Model(user).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Refresh user data
	updatedUser, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	
	updatedUser.Password = ""
	return updatedUser, nil
}

// DeleteUser deletes a user and all associated notes
func (s *UserService) DeleteUser(id string) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	// Delete associated notes first
	if err := database.DB.Where("user_id = ?", id).Delete(&models.Note{}).Error; err != nil {
		return errors.New("error al eliminar notas del usuario")
	}

	// Delete user
	if err := database.DB.Delete(user).Error; err != nil {
		return errors.New("error al eliminar usuario")
	}

	return nil
}

// AuthenticateUser validates user credentials
func (s *UserService) AuthenticateUser(req *models.LoginRequest) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("credenciales inválidas")
		}
		return nil, err
	}

	// Check if user is active
	if user.Status != "activo" {
		return nil, errors.New("cuenta inactiva")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Clear password from response
	user.Password = ""
	return &user, nil
}