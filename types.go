package tictail

import (
	"time"
)

// https://tictail.com/developers/documentation/api-reference/#errors
type TError struct {
	Message      string            `json:"message"`
	Status       int               `json:"status"`
	SupportEmail string            `json:"support_email"`
	Params       map[string]string `json:"params"`
}

// Resources

// Store
// https://tictail.com/developers/documentation/api-reference/#store
type TStore struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	Currency         string    `json:"currency"`
	Language         string    `json:"language"`
	Url              string    `json:"url"`
	DashboardUrl     string    `json:"dashboard_url"`
	StorekeeperEmail string    `json:"storekeeper_email"`
	Sandbox          bool      `json:"sandbox"`
	CreatedAt        time.Time `json:"created_at"`
	ModifiedAt       time.Time `json:"modified_at"`
}
