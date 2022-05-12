// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"accountID"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	ToAccountID   int64 `json:"toAccountID"`
	FromAccountID int64 `json:"fromAccountID"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}
