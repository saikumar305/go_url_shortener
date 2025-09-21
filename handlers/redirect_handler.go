package handlers

import (
	"go_url_shortener/models"
	"net/http"
)



func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Path[len("/"):]

	var url models.URL
	result := models.DB.Where("code = ?", code).First(&url)
	if result.Error != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.Original, http.StatusFound)
}