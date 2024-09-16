package main

import (
	"fmt"
	"log"

	"vanminton-be/api/routes"
	"vanminton-be/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load database connection details (e.g., from environment variables or a config file)
	dbConnString := "your_postgres_connection_string_here"
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize database connection
	db, err := database.NewRDSDB(dbConnString)
	fmt.Println(db)
	fmt.Println(err)
	// db, err := database.NewRDSDB(dbConnString)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.Close() // Ensure that the database connection is closed when the main function exits

	// // Initialize the Gin router
	router := gin.Default()

	// // Setup routes
	routes.RegisterRoutes(router)

	// // Register handlers
	// handlers.SetupHandlers(router)

	// // Set the server parameters
	// server := &http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      router,
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	// // Start the server
	// fmt.Println("Starting server on port 8080...")
	// if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 	log.Fatalf("Error starting server: %v", err)
	// }
}

// // createUser is an HTTP handler function that responds to POST requests to create a new user.
// func createUser(c *gin.Context) {
// 	// Declare a new instance of the User struct to store the values from the request body.
// 	var user User

// 	// Attempt to bind the JSON in the request body to the user variable.
// 	// ShouldBindJSON is a Gin method that unmarshals the body of the request into the provided struct based on the 'Content-Type' header, which should be 'application/json' in this case.
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		// If there's an error during the binding process (such as JSON being in the wrong format, or missing required fields),
// 		// respond with a 400 Bad Request status code and include the error message in the response.
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// End the function prematurely to prevent further processing since there was an error.
// 		return
// 	}

// 	// Define the SQL statement that will be used to insert the user data into the database.
// 	// $1, $2, $3 are placeholders for the values that will be safely substituted into the query to prevent SQL injection.
// 	query := `INSERT INTO users (id, firstname, lastname) VALUES ($1, $2, $3)`

// 	// Execute the SQL statement against the database with the values from the user struct.
// 	// db.Exec executes the SQL command with the given parameters and returns a Result object and an error.
// 	// The Result object contains information about the operation, such as how many rows were affected.
// 	// The error object will be non-nil if an error occurred during the query execution.
// 	_, err := db.Exec(query, user.ID, user.FirstName, user.LastName)
// 	// Check if there was an error when executing the SQL statement.
// 	if err != nil {
// 		// If an error occurred (such as a database connection issue, or SQL syntax error),
// 		// respond with a 500 Internal Server Error status code and include the error message in the response.
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		// End the function prematurely to prevent further processing since there was an error.
// 		return
// 	}

// 	// If the insert operation was successful, respond with a 201 Created status code.
// 	// The gin.H is a shortcut in Gin to create a map[string]interface{}, which Gin uses to return a JSON response.
// 	c.JSON(http.StatusCreated, gin.H{"status": "user created"})
// }
