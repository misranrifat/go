package service

import (
	"time"

	"github.com/yourusername/go-crud-api/internal/model"
	"github.com/yourusername/go-crud-api/internal/repository"
)

// UserService defines the interface for user business logic
type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id string) (model.User, error)
	Create(req model.UserCreateRequest) (model.User, error)
	Update(id string, req model.UserUpdateRequest) (model.User, error)
	Delete(id string) error
}

// userService is an implementation of UserService
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// GetAll returns all users
func (s *userService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

// GetByID returns a user by ID
func (s *userService) GetByID(id string) (model.User, error) {
	return s.repo.GetByID(id)
}

// Create creates a new user
func (s *userService) Create(req model.UserCreateRequest) (model.User, error) {
	user := model.NewUser(req.FirstName, req.LastName, req.Email)
	return s.repo.Create(user)
}

// Update updates a user
func (s *userService) Update(id string, req model.UserUpdateRequest) (model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return model.User{}, err
	}

	// Update fields if provided
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	user.UpdatedAt = time.Now()
	return s.repo.Update(id, user)
}

// Delete removes a user
func (s *userService) Delete(id string) error {
	return s.repo.Delete(id)
}
