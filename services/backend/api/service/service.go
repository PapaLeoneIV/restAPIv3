package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"students/db"
	"time"
	"io/ioutil"
	"bytes"
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
	Id     int32  `json:"id"`
	Name	string  `json:"name"`
	Subject  string  `json:"subject"`
	Body 	string  `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



func (s *Service) CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateProduct Handler called")
	if r.Method == http.MethodOptions {
		fmt.Println("Options request!")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
	fmt.Println("Post request!")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	ctx := context.Background()

	var message db.CreateProductParams
/* 	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	bodyString := string(bodyBytes)
	fmt.Printf("Request body: %s\n", bodyString)
		if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}
	fmt.Printf("Decoded message: %+v\n", message) */

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max memory
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("Error parsing request body: %v\n", err)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("Error parsing request body: %v\n", err)
		return
	}

	
	
 	message.Name = r.FormValue("name")
	message.Subject = r.FormValue("subject")
	message.Body = r.FormValue("body")
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	
	if lastId, err := s.db.GetLastIdx(ctx); err != nil {
		http.Error(w, "Failed to fetch message", http.StatusInternalServerError)
		fmt.Printf("Error fetching message: %v\n", err)
		return
	} else {
		fmt.Printf("Last ID: %d\n", lastId)
		message.Id = lastId + 1
		fmt.Printf("Last ID: %d\n", message.Id)

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
	if id2 , err := strconv.Atoi(id); err != nil {
		http.Error(w, "Invalid user ID in context", http.StatusBadRequest)
		fmt.Println("Invalid user ID in context")
		return
	} else {
		out, err := s.db.GetProduct(ctx, int32(id2))
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
}

func (s *Service) ListAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListAllProducts Handler called")
	ctx := context.Background()
	fmt.Printf("Context: %+v\n", ctx)
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

	if id2 , err := strconv.Atoi(id); err != nil {
		http.Error(w, "Invalid user ID in context", http.StatusBadRequest)
		fmt.Println("Invalid user ID in context")
		return
	} else {
		err := s.db.DeleteProduct(ctx, int32(id2))
		if err != nil {
			http.Error(w, "Failed to fetch message", http.StatusInternalServerError)
			fmt.Printf("Error fetching message: %v\n", err)
			return
		}
		fmt.Println("Message deleted successfully")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message deleted successfully"))
	}
}  