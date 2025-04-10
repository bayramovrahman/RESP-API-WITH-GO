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
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`

	_, err := DB.Exec(createEventsTable) // SQL sorgusu çalıştırılıyor

	if err != nil {
		panic("Could not create events table: " + err.Error()) // Tablo oluşturulamazsa hata fırlat
	}
}
