package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/yourusername/go-crud-api/internal/model"
)

// Common errors
var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user with this email already exists")
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id string) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(id string, user model.User) (model.User, error)
	Delete(id string) error
}

// userRepository is an in-memory implementation of UserRepository
type userRepository struct {
	users map[string]model.User
	mutex sync.RWMutex
}

// NewUserRepository creates a new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		users: make(map[string]model.User),
	}
}

// GetAll returns all users
func (r *userRepository) GetAll() ([]model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]model.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

// GetByID returns a user by ID
func (r *userRepository) GetByID(id string) (model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return model.User{}, ErrUserNotFound
	}
	return user, nil
}

// Create adds a new user
func (r *userRepository) Create(user model.User) (model.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if user with this email already exists
	for _, existingUser := range r.users {
		if existingUser.Email == user.Email {
			return model.User{}, ErrUserExists
		}
	}

	r.users[user.ID] = user
	return user, nil
}

// Update updates an existing user
func (r *userRepository) Update(id string, user model.User) (model.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.users[id]
	if !exists {
		return model.User{}, ErrUserNotFound
	}

	// Check if email is already taken by another user
	for userId, existingUser := range r.users {
		if userId != id && existingUser.Email == user.Email {
			return model.User{}, ErrUserExists
		}
	}

	// Update timestamp
	user.UpdatedAt = time.Now()
	r.users[id] = user
	return user, nil
}

// Delete removes a user
func (r *userRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.users[id]
	if !exists {
		return ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}
