package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const contextKey = "id"

type Service struct {
	db *sql.DB
}

func NewService(d *sql.DB) *Service {
	fmt.Printf("Service handler created\n")
	return &Service{db: d}
}

type Message struct {
	Id     string  `json:"id"`
	Name	string  `json:"name"`
	Subject  string  `json:"subject"`
	Body 	string  `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



func (s *Service) CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateProduct Handler called")
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}
	fmt.Printf("Decoded message: %+v\n", message)

	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	query := `INSERT INTO students (id, name, subject, body, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Exec(query, message.Id, message.Name, message.Subject, message.Body, message.CreatedAt, message.UpdatedAt)
	if err != nil {
		http.Error(w, "Failed to insert message", http.StatusInternalServerError)
		fmt.Printf("Error inserting message into database: %v\n", err)
		return
	}
	fmt.Println("message inserted successfully")

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Service) GetProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetProduct Handler called")
	ctx := r.Context()
	id, ok := ctx.Value(contextKey).(string)
	if id == "" || !ok {
		http.Error(w, "Invalid or missing user ID in context", http.StatusBadRequest)
		fmt.Println("Invalid or missing user ID in context")
		return
	}
	fmt.Printf("Context ID: %s\n", id)

	var album Message
	err := s.db.QueryRowContext(ctx, "SELECT * FROM students WHERE id=$1", id).Scan(
			&album.Id, &album.Name, &album.Subject, &album.Body, &album.CreatedAt, &album.UpdatedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "Message not found", http.StatusNotFound)
		fmt.Printf("Message not found for ID: %s\n", id)
		return
	} else if err != nil {
		http.Error(w, "Error fetching message", http.StatusInternalServerError)
		fmt.Printf("Error fetching message: %v\n", err)
		return
	}
	fmt.Printf("Fetched message: %+v\n", album)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(album); err != nil {
		fmt.Printf("Failed to encode message to JSON: %v\n", err)
		http.Error(w, "Failed to encode message to JSON", http.StatusInternalServerError)
	}
}
 
func (s *Service) ListAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListAllProducts Handler called")
	ctx := r.Context()

	rows, err := s.db.QueryContext(ctx, "SELECT * FROM students")
	if err == sql.ErrNoRows {
		http.Error(w, "No message inserted in the database yett", http.StatusNotFound)
		fmt.Printf("No message to be retrieved\n")
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch albums", http.StatusInternalServerError)
		fmt.Printf("Error fetching albums: %v\n", err)
		return
	}
	defer rows.Close()
	fmt.Println("Messages fetched from database")

	var messages []Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.Id, &message.Name, &message.Subject, &message.Body,&message.CreatedAt, &message.UpdatedAt); err != nil {
			http.Error(w, "Error scanning message row", http.StatusInternalServerError)
			fmt.Printf("Error scanning message row: %v\n", err)
			return
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating message rows", http.StatusInternalServerError)
		fmt.Printf("Error iterating message rows: %v\n", err)
		return
	}
	fmt.Printf("Fetched messages: %+v\n", messages)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		fmt.Printf("Failed to encode messages to JSON: %v\n", err)
		http.Error(w, "Failed to encode messages to JSON", http.StatusInternalServerError)
	}
}

func (s *Service) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateProduct Handler called")

	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}
	fmt.Printf("Decoded message for update: %+v\n", message)
	
	message.UpdatedAt = time.Now()
	dbQuery := `UPDATE students SET name=$1, subject=$2, body=$3, created_at=$4, updated_at=$5 WHERE id=$6`
	effected, err := s.db.Exec(dbQuery, message.Name, message.Subject, message.Body,message.CreatedAt, message.UpdatedAt, message.Id)
	if err != nil {
		http.Error(w, "Failed to update message", http.StatusInternalServerError)
		fmt.Printf("Error updating message in database: %v\n", err)
		return
	}
	if rows, _ := effected.RowsAffected(); rows == 0 {
		http.Error(w, "message not found", http.StatusNotFound)
		fmt.Printf("message not found for ID: %s\n", message.Id)
		return
	}
	
	fmt.Println("Message updated successfully")
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Service) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteProduct Handler called")
	ctx := r.Context()
	id, ok := ctx.Value(contextKey).(string)
	if id == "" || !ok {
		http.Error(w, "Invalid or missing user ID in context", http.StatusBadRequest)
		fmt.Println("Invalid or missing user ID in context")
		return
	}
	fmt.Printf("Context ID for delete: %s\n", id)

	dbQuery := `DELETE FROM students WHERE id=$1`
	effected, err := s.db.Exec(dbQuery, id)
	if rows, _ := effected.RowsAffected(); rows == 0 {
		http.Error(w, "Message not found", http.StatusNotFound)
		fmt.Printf("Message not found for ID: %s\n", id)
		return
	}
	if err != nil {
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		fmt.Printf("Error deleting message from database: %v\n", err)
		return
	}
	fmt.Println("Message deleted successfully")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message deleted successfully"))
} 