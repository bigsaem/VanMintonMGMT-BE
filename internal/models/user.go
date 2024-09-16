package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID               string         `json:"id"`
	FirstName        string         `json:"first_name"`
	LastName         string         `json:"last_name"`
	Email            sql.NullString `json:"email"`
	Phone            sql.NullString `json:"phone"`
	CourtCredit      float64        `json:"court_credit"`
	MembershipCredit float64        `json:"membership_credit"`
	LessonCredit     float64        `json:"lesson_credit"`
	SignupDate       time.Time      `json:"signup_date"`
	ExpiryDate       time.Time      `json:"expiry_date"`
	Status           string         `json:"status"`
	Role             string         `json:"role"`
}
