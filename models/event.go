package models

import (
	"time"
	"example.com/rest-api/database"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query) // SQL sorgusu çalıştırılmadan önce hazırlanıyor

	if err != nil {
		return err 
	}
	defer stmt.Close() // Fonksiyon bitiminde statement kapatılıyor

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	// Hazırlanan sorgu çalıştırılıyor ve ilgili alanlara Event nesnesindeki veriler yerleştiriliyor

	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // Eklenen verinin otomatik oluşturulan ID’si alınıyor

	e.ID = id // Bu ID, Event nesnesine atanıyor (yalnız burada değer kopyalanmış, dışarı yansımaz)

	return err

	// events = append(events, e)
}


func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query) // Sorgu çalıştırılıyor ve sonuçlar alınmaya başlanıyor

	if err != nil {
		return nil, err
	}
	defer rows.Close() // rows işlemi bittiğinde kapatılıyor

	var events []Event // Event tipinde slice tanımlanıyor

	for rows.Next() { // Her satır için dön
		var event Event // Yeni bir Event nesnesi oluşturuluyor

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // Satırdaki veriler Event nesnesine aktarılıyor

		if err != nil {
			return nil, err // Eğer hata varsa işlemi sonlandır ve hatayı döndür
		}

		events = append(events, event) // Event nesnesi listeye ekleniyor
	}

	return events, nil // Event listesi ve hata bilgisi döndürülüyor (hata yoksa nil olur)
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id) // Veritabanında tek bir satır döndüren sorguyu çalıştırır.

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // Satırdaki veriler Event nesnesine aktarılıyor
	if err!= nil {
		return nil, err // Eğer hata varsa işlemi sonlandır ve hatayı döndür
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query) // SQL sorgusu çalıştırılmadan önce hazırlanıyor

	if err != nil {
		return err 
	}
	defer stmt.Close() // Fonksiyon bitiminde statement kapatılıyor

	// Hazırlanan sorguyu çalıştır ve parametreleri ver
	// event struct'ındaki alanları sorguya parametre olarak geçiyoruz
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	// Hata oluştuysa hatayı döndür, yoksa nil dönecek (başarılı)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query) // SQL sorgusu çalıştırılmadan önce hazırlanıyor

	if err != nil {
		return err 
	}
	defer stmt.Close() // Fonksiyon bitiminde statement kapatılıyor

	_, err = stmt.Exec(event.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query) // SQL sorgusu çalıştırılmadan önce hazırlanıyor
	if err != nil {
		return err 
	}
	defer stmt.Close() // Fonksiyon bitiminde statement kapatılıyor

	_, err = stmt.Exec(e.ID, e.UserID)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query) // SQL sorgusu çalıştırılmadan önce hazırlanıyor

	if err != nil {
		return err 
	}
	defer stmt.Close() // Fonksiyon bitiminde statement kapatılıyor

	_, err = stmt.Exec(e.ID, userId)
	return err
}
