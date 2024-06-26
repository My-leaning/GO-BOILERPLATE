package postgresDB

import (
	"database/sql"
	"fmt"
	"go_boilerplate/internal/util"
	"log"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NOTE this function is used to connect to the postgres database (Don't use Type orm)
// func PostgresConnect() {
// 	// Define the PostgreSQL connection string
// 	connStr := "user=postgres dbname=test password=12345678 host=localhost sslmode=disable"

// 	// Open a connection to the database
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatalf("Failed to open the database: %v", err)
// 	}

// 	// Close the connection when the main function finishes
// 	defer db.Close()

// 	// Ping the database to verify the connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	fmt.Println("Successfully connected to the database!")
// }

// NOTE this function is used to connect to the postgres database (use Type orm)
var DB *gorm.DB

func PostgresConnect() {
	// Load configuration
	config, err := util.LoadConfig("./")

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connection string to connect to the PostgreSQL server without specifying a database
	serverDsn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable", config.DBuser, config.DBpassword, config.DBhost, config.DBport)

	// Open a connection to the PostgreSQL server
	serverDb, err := sql.Open("postgres", serverDsn)
	if err != nil {
		log.Fatalf("Failed to connect to the PostgreSQL server: %v", err)
	}
	defer serverDb.Close()

	// Check if the database exists
	var exists bool
	checkDbQuery := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", config.DBname)
	err = serverDb.QueryRow(checkDbQuery).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}

	// Create the database if it doesn't exist
	if !exists {
		createDbQuery := fmt.Sprintf("CREATE DATABASE %s", config.DBname)
		_, err = serverDb.Exec(createDbQuery)
		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		log.Printf("Database %s created successfully", config.DBname)
	}

	// Define the PostgreSQL connection string for the specific database
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", config.DBuser, config.DBname, config.DBpassword, config.DBhost, config.DBport)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Admin{})

	// Assign the DB variable
	DB = db
}

type User struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string     `gorm:"unique;not null" json:"username"`
	Password  string     `gorm:"not null" json:"password"`
	Phone     string     `gorm:"unique;not null" json:"phone"`
	CreatedAt time.Time  `gorm:"autoCreateTime"  json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type Admin struct {
	ID       primitive.ObjectID `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password,omitempty"`
	Phone    string             `json:"phone,omitempty"`
}
