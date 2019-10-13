package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetNews(w http.ResponseWriter, r *http.Request) {
	newsId := r.FormValue("id")

	id, err := strconv.Atoi(newsId)
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, "Failed to parse ID")
		return
	}

	storage := StorageClient{}
	err = storage.Connect("amqp://guest:guest@localhost:5672/") // TODO: Get URL from CLI
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	defer storage.Close()

	response, err := storage.GetNews(id)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to get News")
		return
	}

	out, err := json.Marshal(response)
	_, _ = w.Write(out)
}

func AddNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	title := r.FormValue("title")
	date := r.FormValue("date")

	_, err := time.Parse(time.RFC822Z, date)
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, "Failed to parse date")
		return
	}

	storage := StorageClient{}
	err = storage.Connect("amqp://guest:guest@localhost:5672/") // TODO: Get URL from CLI
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	defer storage.Close()

	response, err := storage.AddNews(title, date)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to add News")
		return
	}

	out, err := json.Marshal(response)
	_, _ = w.Write(out)
}

func writeJsonResponse(w http.ResponseWriter, status int, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
	}
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func writeJsonError(w http.ResponseWriter, status int, msg string) {
	log.Printf("Error: %s", msg)
	writeJsonResponse(w, status, ErrorResponse{Code: status, Message: msg})
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AddNewsResponse struct {
	Id uint32 `json:"id"`
}

type GetNewsResponse struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}
