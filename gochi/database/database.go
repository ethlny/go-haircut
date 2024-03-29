package database

import (
	"database/sql"
	"log"
	"os"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SetupDatabase() *sql.DB {
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	createTables(db)
	return db
}

func ShowTables(db *sql.DB) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		log.Println(tableName)
	}
}

func checkError(err error, context string) {
	if err != nil {
		log.Fatalf("failed to %s: %v", context, err)
	}
}

func createTables(db *sql.DB) {
	// Client table
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS clients (
				uid VARCHAR(255) PRIMARY KEY,
				email VARCHAR(255) NOT NULL,
				age INT NOT NULL,
				password VARCHAR(255) NOT NULL
			)
			`)
	checkError(err, "clients")

	// Hairdresser table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS hairdressers (
				uid VARCHAR(255) PRIMARY KEY,
				saloonID VARCHAR(255) NOT NULL,
				firstName VARCHAR(255) NOT NULL,
				speciality VARCHAR(255) NOT NULL
				)
				`)
	checkError(err, "hairdressers")

	// Admin table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS admin (
				uid VARCHAR(255) PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL
				)
				`)
	checkError(err, "admin")

	// Hair saloon table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS hairSaloon (
				uid VARCHAR(255) PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				address VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				phone VARCHAR(255) NOT NULL,
				openingTime TIME NOT NULL,
				closingTime TIME NOT NULL,
				password VARCHAR(255) NOT NULL
				)`)
	checkError(err, "HairSaloon")

	// Appointments table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS appointments (
				uid VARCHAR(255) PRIMARY KEY,
				saloonID VARCHAR(255) NOT NULL,
				clientID VARCHAR(255) NOT NULL,
				startHour TIME NOT NULL,
				hairdresserID VARCHAR(255) NULL,
				status VARCHAR(255) NOT NULL,
				appointmentDate DATE
				)
									`)
	checkError(err, "appointments")

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var EmailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func CheckEmail(email string) bool {
	return EmailRegex.MatchString(email)
}
