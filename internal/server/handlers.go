package server

import (
	"Lnkio/internal/shorturl"
	"encoding/json"
	"fmt"
	"net/http"
)

type UrlResponse struct {
	Url string `json:"url"`
}

func (s *Server) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if s.env == "local" {
		url = fmt.Sprintf("https://localhost:8080%s", url)
	}
	fmt.Println(url)
	repo := shorturl.NewUrlRepository(s.db.DB())
	baseUrl, err := repo.GetLink(url)
	fmt.Println(baseUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, baseUrl, http.StatusMovedPermanently)
}

func (s *Server) getUrlsHandler(w http.ResponseWriter, r *http.Request) {
	repo := shorturl.NewUrlRepository(s.db.DB())
	urls, err := repo.GetLinks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(urls)
}

func (s *Server) newUrlHandler(w http.ResponseWriter, r *http.Request) {
	var u shorturl.Url

	r.Body = http.MaxBytesReader(w, r.Body, 1<<10)

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if s.env == "local" {
		u.ShortUrl = fmt.Sprintf("https://localhost:8080/%s", u.ShortUrl)
	}

	if err := u.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := shorturl.NewUrlRepository(s.db.DB())
	url, err := repo.CreateLink(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UrlResponse{Url: url.ShortUrl})
}
