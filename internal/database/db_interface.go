package database

import "vanminton-be/internal/models"

// Database represents the interface for RDS database operations.
type Database interface {
	CreateUser(user models.User) error
	GetUser(id string) (models.User, error)
	// Add other necessary methods
}
