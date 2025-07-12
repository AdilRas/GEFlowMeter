package clicker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Clicker represents the API client for the Clicker service
type Clicker struct {
	baseURL    string
	token      string
	httpClient *http.Client
	user       *User
}

// New creates a new Clicker client instance
func New() *Clicker {
	return &Clicker{
		baseURL:    "https://clicker.iiindia.org",
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// Login authenticates with the Clicker API using provided credentials
func (c *Clicker) Login(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password cannot be empty")
	}

	loginReq := LoginRequest{
		Username: username,
		Password: password,
	}

	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		return fmt.Errorf("failed to marshal login request: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/api/auth/login", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create login request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute login request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read login response: %w", err)
	}

	var loginResp LoginResponse
	if err := json.Unmarshal(body, &loginResp); err != nil {
		return fmt.Errorf("failed to parse login response: %w", err)
	}

	if !loginResp.Success {
		return errors.New("login failed: API returned success=false")
	}

	c.token = loginResp.Token
	c.user = &loginResp.User

	return nil
}

// GetAreas fetches the areas data from the Clicker API
func (c *Clicker) GetAreas() (*AreasResponse, error) {
	if c.token == "" {
		return nil, errors.New("not logged in: call Login() first")
	}

	req, err := http.NewRequest("GET", c.baseURL+"/api/proxy/areas", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create areas request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "*/*")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute areas request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("areas request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read areas response: %w", err)
	}

	var areasResp AreasResponse
	if err := json.Unmarshal(body, &areasResp); err != nil {
		return nil, fmt.Errorf("failed to parse areas response: %w", err)
	}

	return &areasResp, nil
}

// IsLoggedIn returns true if the client is currently authenticated
func (c *Clicker) IsLoggedIn() bool {
	return c.token != ""
}

// GetUser returns the current user information if logged in
func (c *Clicker) GetUser() (*User, error) {
	if c.token == "" {
		return nil, errors.New("not logged in: call Login() first")
	}
	return c.user, nil
}

// GetToken returns the current authentication token if logged in
func (c *Clicker) GetToken() (string, error) {
	if c.token == "" {
		return "", errors.New("not logged in: call Login() first")
	}
	return c.token, nil
}
