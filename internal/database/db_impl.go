package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"vanminton-be/internal/models"

	_ "github.com/lib/pq"
)

// ...

// RDSDB is an implementation of the Database interface for Amazon RDS with PostgreSQL.
type RDSDB struct {
	db *sql.DB
}

// NewRDSDB creates a new connection to the RDS PostgreSQL instance.
func NewRDSDB(dataSourceName string) (string, error) {
	dbHost := os.Getenv("DB_URL")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	// region := "us-east-2"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		dbHost, dbPort, dbUser, os.Getenv("DB_PASSWORD"), dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("hi")
		log.Fatal(err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Println("hi")
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
	result, err := db.Query((`SELECT * from "User";`))
	if err != nil {
		panic(err.Error())
	}
	log.Println((result))
	defer result.Close()

	var users []models.User
	for result.Next() {
		var user models.User
		err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.CourtCredit, &user.MembershipCredit, &user.LessonCredit, &user.SignupDate, &user.ExpiryDate, &user.Status, &user.Role)
		if err != nil {
			log.Println(`something happened`)
			log.Fatal(err)
		}
		users = append(users, user)
	}

	for _, user := range users {
		fmt.Printf("ID: %s, FirstName: %s, LastName: %s\n", user.ID, user.FirstName, user.LastName)
	}

	// cfg, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	panic("configuration error: " + err.Error())
	// }
	// dbEndpoint := "vanminton-db.cu4yoqsaub61.us-east-2.rds.amazonaws.com:2964"
	// dbName := "VanMinton"
	// dbUser := "VanmintonAdmin"
	// region := "us-east-2"
	// authenticationToken, err := auth.BuildAuthToken(
	// 	context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)
	// if err != nil {
	// 	panic("failed to create authentication token: " + err.Error())
	// }
	// dsn := fmt.Sprintf("endpoint=%s user=%s password=%s dbname=%s",
	// 	dbEndpoint, dbUser, authenticationToken, dbName,
	// )

	// db, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	return "connected", err
}

// func (rds *RDSDB) CreateUser(user models.User) error {
// 	// Implement user creation in RDS PostgreSQL
// 	// Example: INSERT INTO users (columns...) VALUES (values...)
// 	// ...
// }

// func (rds *RDSDB) GetUser(id string) (models.User, error) {
// 	// Implement retrieval of a user from RDS PostgreSQL
// 	// Example: SELECT * FROM users WHERE id = $1
// 	// ...
// }
