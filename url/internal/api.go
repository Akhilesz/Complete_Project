package internal

import (
	"encoding/json"
	"net/http"
	"url/repository"
)

type Handler struct {
	Repo *repository.URLRepo
}

func (handler *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	shortCode := GenerateShortCode(req.LongURL)
	code, err := handler.Repo.Save(req.LongURL, shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]string{"short_url": "http://localhost:8080/" + code})
}

func (handler *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	longURL, err := handler.Repo.Get(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}
