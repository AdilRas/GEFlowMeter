package clicker

import "time"

// LoginRequest represents the request payload for the login API
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User represents the user object in the login response
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginResponse represents the response from the login API
type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	User    User   `json:"user"`
}

// Area represents an individual area object
type Area struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Capacity     int       `json:"capacity"`
	CurrentCount int       `json:"current_count"`
	Status       string    `json:"status"`
	IsEnabled    bool      `json:"isEnabled"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// AreasResponse represents the response from the areas API
type AreasResponse struct {
	Areas   []Area `json:"areas"`
	Success bool   `json:"success"`
}
