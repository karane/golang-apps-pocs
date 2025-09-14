package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *service.URLService
}

func NewHandler(s *service.URLService) *Handler {
	return &Handler{Service: s}
}


func (h *Handler) CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OriginalURL string `json:"original_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url, err := h.Service.CreateShortURL(req.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(url)
}

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	url, err := h.Service.GetOriginalURL(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func (h *Handler) Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", h.CreateShortURLHandler).Methods("POST")
	r.HandleFunc("/r/{code}", h.RedirectHandler).Methods("GET")
	return r
}
