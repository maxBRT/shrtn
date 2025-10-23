package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shrtn/internal/shorturl"
)

type UrlResponse struct {
	Url string `json:"url"`
}

func (s *Server) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	switch url {
	case "/":
		http.ServeFile(w, r, "./public/index.html")
	default:
		url = fmt.Sprintf("https://shrtn.it.com%s", url)
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

	u.ShortUrl = fmt.Sprintf("https://shrtn.it.com/%s", u.ShortUrl)

	if err := u.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := shorturl.NewUrlRepository(s.db.DB())
	url, err := repo.CreateLink(u)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UrlResponse{Url: url.ShortUrl})
}
