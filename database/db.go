package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // "api.db" adlı SQLite veritabanını açıyoruz

	if err != nil {
		panic("Couldn't connect to database") // Bağlantı başarısızsa program durduruluyor
	}

	DB.SetMaxOpenConns(10)    // Aynı anda açılabilecek maksimum bağlantı sayısı 10
	DB.SetMaxIdleConns(5)     // Boşta bekleyen maksimum bağlantı sayısı 5

	createTables()            // Gerekli tablolar oluşturuluyor
}

func createTables() {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create users table: " + err.Error()) // Tablo oluşturulamazsa hata fırlat
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id) 
		)
	`

	_, err = DB.Exec(createEventsTable) // SQL sorgusu çalıştırılıyor
	if err != nil {
		panic("Could not create events table: " + err.Error()) // Tablo oluşturulamazsa hata fırlat
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createRegistrationsTable) // SQL sorgusu çalıştırılıyor
	if err != nil {
		panic("Could not create registrations table: " + err.Error()) // Tablo oluşturulamazsa hata fırlat
	}
}
