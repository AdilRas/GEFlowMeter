package main

import (
	"fmt"
	"log"
	"os"

	"ge-flow-meter/clicker"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Clicker API Client Demo")

	// Try to load .env.local file - ignore error if file doesn't exist
	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Println("No .env.local file found - using system environment variables")
	} else {
		fmt.Println("Loaded .env.local file")
	}

	// Create a new clicker client
	client := clicker.New()

	// Check if already logged in
	if !client.IsLoggedIn() {
		fmt.Println("Not logged in. Attempting to login...")

		// Get credentials from environment variables
		username := os.Getenv("CLICKER_USERNAME")
		password := os.Getenv("CLICKER_PASSWORD")

		if username == "" || password == "" {
			log.Fatal("CLICKER_USERNAME and CLICKER_PASSWORD environment variables must be set")
		}

		// Login using credentials
		if err := client.Login(username, password); err != nil {
			log.Fatalf("Login failed: %v", err)
		}

		fmt.Println("Login successful!")
	}

	// Get user information
	user, err := client.GetUser()
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	fmt.Printf("Logged in as: %s (ID: %d, Role: %s)\n", user.Username, user.ID, user.Role)

	// Get areas data
	areas, err := client.GetAreas()
	if err != nil {
		log.Fatalf("Failed to get areas: %v", err)
	}

	fmt.Printf("\nFound %d areas:\n", len(areas.Areas))
	for _, area := range areas.Areas {
		fmt.Printf("- %s: %d/%d capacity (Status: %s)\n",
			area.Name, area.CurrentCount, area.Capacity, area.Status)
	}
}
