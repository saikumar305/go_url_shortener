package handlers

import (
	"encoding/json"
	"fmt"
	"go_url_shortener/models"
	"math/rand"
	"net/http"
)




type ShortenRequest struct {
	URL string `json:"url"`
}


type ShortenResponse struct {
	ShortUrl string `json:"short_url"`
}

func generateShortCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req);
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	code := generateShortCode(6)

	shortUrl := fmt.Sprintf("http://localhost:8080/%s", code)

	url := models.URL{
		Code:      code,
		Original:  req.URL,
		ShortUrl:  shortUrl,
	}
	result := models.DB.Create(&url)
	if result.Error != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	resp := ShortenResponse{ShortUrl: shortUrl}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

