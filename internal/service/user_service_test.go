package service

import (
	"errors"
	"testing"

	"github.com/yourusername/go-crud-api/internal/model"
	"github.com/yourusername/go-crud-api/internal/repository"
)

// mockUserRepository is a mock implementation of the repository.UserRepository interface
type mockUserRepository struct {
	users map[string]model.User
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: make(map[string]model.User),
	}
}

func (m *mockUserRepository) GetAll() ([]model.User, error) {
	users := make([]model.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

func (m *mockUserRepository) GetByID(id string) (model.User, error) {
	user, exists := m.users[id]
	if !exists {
		return model.User{}, repository.ErrUserNotFound
	}
	return user, nil
}

func (m *mockUserRepository) Create(user model.User) (model.User, error) {
	for _, existingUser := range m.users {
		if existingUser.Email == user.Email {
			return model.User{}, repository.ErrUserExists
		}
	}
	m.users[user.ID] = user
	return user, nil
}

func (m *mockUserRepository) Update(id string, user model.User) (model.User, error) {
	_, exists := m.users[id]
	if !exists {
		return model.User{}, repository.ErrUserNotFound
	}

	// Check if email is already taken by another user
	for userID, existingUser := range m.users {
		if userID != id && existingUser.Email == user.Email {
			return model.User{}, repository.ErrUserExists
		}
	}

	m.users[id] = user
	return user, nil
}

func (m *mockUserRepository) Delete(id string) error {
	_, exists := m.users[id]
	if !exists {
		return repository.ErrUserNotFound
	}
	delete(m.users, id)
	return nil
}

func TestUserService_Create(t *testing.T) {
	mockRepo := newMockUserRepository()
	userService := NewUserService(mockRepo)

	testCases := []struct {
		name          string
		request       model.UserCreateRequest
		expectedError error
	}{
		{
			name: "Valid user creation",
			request: model.UserCreateRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
			},
			expectedError: nil,
		},
		{
			name: "Duplicate email",
			request: model.UserCreateRequest{
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "john.doe@example.com", // Same as above
			},
			expectedError: repository.ErrUserExists,
		},
	}

	// First create a valid user
	_, err := userService.Create(testCases[0].request)
	if err != nil {
		t.Fatalf("Failed to create initial user: %v", err)
	}

	// Then try to create a duplicate
	_, err = userService.Create(testCases[1].request)
	if !errors.Is(err, testCases[1].expectedError) {
		t.Errorf("Expected error %v, got %v", testCases[1].expectedError, err)
	}
}

func TestUserService_GetByID(t *testing.T) {
	mockRepo := newMockUserRepository()
	userService := NewUserService(mockRepo)

	// Create a test user
	createReq := model.UserCreateRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	createdUser, err := userService.Create(createReq)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Test getting existing user
	user, err := userService.GetByID(createdUser.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user.ID != createdUser.ID {
		t.Errorf("Expected user ID %s, got %s", createdUser.ID, user.ID)
	}

	// Test getting non-existent user
	_, err = userService.GetByID("non-existent-id")
	if !errors.Is(err, repository.ErrUserNotFound) {
		t.Errorf("Expected error %v, got %v", repository.ErrUserNotFound, err)
	}
}

func TestUserService_Update(t *testing.T) {
	mockRepo := newMockUserRepository()
	userService := NewUserService(mockRepo)

	// Create a test user
	createReq := model.UserCreateRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	createdUser, err := userService.Create(createReq)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Test updating existing user
	updateReq := model.UserUpdateRequest{
		FirstName: "Johnny",
		LastName:  "Doe",
		Email:     "johnny.doe@example.com",
	}
	updatedUser, err := userService.Update(createdUser.ID, updateReq)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if updatedUser.FirstName != updateReq.FirstName {
		t.Errorf("Expected first name %s, got %s", updateReq.FirstName, updatedUser.FirstName)
	}
	if updatedUser.Email != updateReq.Email {
		t.Errorf("Expected email %s, got %s", updateReq.Email, updatedUser.Email)
	}

	// Test updating non-existent user
	_, err = userService.Update("non-existent-id", updateReq)
	if !errors.Is(err, repository.ErrUserNotFound) {
		t.Errorf("Expected error %v, got %v", repository.ErrUserNotFound, err)
	}
}

func TestUserService_Delete(t *testing.T) {
	mockRepo := newMockUserRepository()
	userService := NewUserService(mockRepo)

	// Create a test user
	createReq := model.UserCreateRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	createdUser, err := userService.Create(createReq)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Test deleting existing user
	err = userService.Delete(createdUser.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify user was deleted
	_, err = userService.GetByID(createdUser.ID)
	if !errors.Is(err, repository.ErrUserNotFound) {
		t.Errorf("Expected error %v, got %v", repository.ErrUserNotFound, err)
	}

	// Test deleting non-existent user
	err = userService.Delete("non-existent-id")
	if !errors.Is(err, repository.ErrUserNotFound) {
		t.Errorf("Expected error %v, got %v", repository.ErrUserNotFound, err)
	}
}
