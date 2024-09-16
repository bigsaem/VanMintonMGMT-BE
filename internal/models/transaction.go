package models

import (
	"time"
)

type Transaction struct {
	ID      int       `json:"id"`      // The ID, will be automatically generated by the database
	Date    time.Time `json:"date"`    // The date of the transaction
	Value   float64   `json:"value"`   // The value of the transaction
	UserID  string    `json:"userid"`  // The user ID, represented as a string UUID
	Purpose string    `json:"purpose"` // The purpose of the transaction
}
