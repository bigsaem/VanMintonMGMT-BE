package models

type Ledger struct {
	ID           int     `json:"id"`           // The ID, will be automatically generated by the database
	SessionID    int     `json:"sessionid"`    // Foreign key reference to the Session table
	UserID       string  `json:"userid"`       // Foreign key reference to the User table, represented as a string UUID
	MoneyToPay   float64 `json:"moneytopay"`   // The amount of money to be paid
	PayConfirmed bool    `json:"payconfirmed"` // Whether the payment has been confirmed, with a default value of false
}
