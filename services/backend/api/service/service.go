package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"students/db"
	"time"
)

const contextKey = "id"

type Service struct {
	db *db.Queries
}

func NewService(d *db.Queries) *Service {
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

	ctx := context.Background()


	var message db.CreateProductParams
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}
	fmt.Printf("Decoded message: %+v\n", message)

	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()

	out, err := s.db.CreateProduct(ctx, message)
	fmt.Printf("Fetched message: %+v\n", out)
	if err != nil {
		http.Error(w, "Failed to fetch message", http.StatusInternalServerError)
		fmt.Printf("Error fetching message: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Service) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value(contextKey).(string)
	if id == "" || !ok {
		http.Error(w, "Invalid or missing user ID in context", http.StatusBadRequest)
		fmt.Println("Invalid or missing user ID in context")
		return
	}
	fmt.Printf("Context ID: %s\n", id)

	out, err := s.db.GetProduct(ctx, id)
	fmt.Printf("Fetched message: %+v\n", out)
	if err != nil {
		http.Error(w, "Failed to fetch message", http.StatusInternalServerError)
		fmt.Printf("Error fetching message: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(out); err != nil {
		fmt.Printf("Failed to encode message to JSON: %v\n", err)
		http.Error(w, "Failed to encode message to JSON", http.StatusInternalServerError)
	}
}

func (s *Service) ListAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListAllProducts Handler called")
	ctx := r.Context()

	out, err := s.db.GetProducts(ctx)
	fmt.Printf("Fetched message: %+v\n", out)
	if err != nil {
		http.Error(w, "Failed to fetch message", http.StatusInternalServerError)
		fmt.Printf("Error fetching message: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(out); err != nil {
		fmt.Printf("Failed to encode message to JSON: %v\n", err)
		http.Error(w, "Failed to encode message to JSON", http.StatusInternalServerError)
	}
}

func (s *Service) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateProduct Handler called")
	ctx := r.Context()
	var message db.UpdateProductParams
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}
	fmt.Printf("Decoded message for update: %+v\n", message)
	
	message.UpdatedAt = time.Now()

	err := s.db.UpdateProduct(ctx, message)
	if err != nil {
		http.Error(w, "Failed to fetch message", http.StatusInternalServerError)
		fmt.Printf("Error fetching message: %v\n", err)
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


	err := s.db.DeleteProduct(ctx, id)
	if err != nil {
		http.Error(w, "Failed to fetch message", http.StatusInternalServerError)
		fmt.Printf("Error fetching message: %v\n", err)
		return
	}
	fmt.Println("Message deleted successfully")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message deleted successfully"))
}  