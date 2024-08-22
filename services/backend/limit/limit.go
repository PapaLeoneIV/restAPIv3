package limit

import (
	"encoding/json"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func RateLimiter(next http.HandlerFunc) http.Handler {
	type Client struct {
		lastSeen    time.Time
		rateLimiter *rate.Limiter
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*Client)
	)

	// Background goroutine to clean up inactive clients
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > time.Minute*3 {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		mu.Lock()
		client, exists := clients[ip]
		if !exists {
			client = &Client{
				rateLimiter: rate.NewLimiter(2, 4), // Adjust rate limits as needed
				lastSeen:    time.Now(),
			}
			clients[ip] = client
		}
		client.lastSeen = time.Now()

		// Check if the request is allowed
		if !client.rateLimiter.Allow() {
			mu.Unlock()
			message := Message{
				Status: "Request Failed",
				Body:   "The API is at capacity, try again later.",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		}
		mu.Unlock()

		next(w, r)
	})
}
