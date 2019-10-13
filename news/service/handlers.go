package service

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetNews(w http.ResponseWriter, r *http.Request) {
	newsId := r.FormValue("id")

	id, err := strconv.Atoi(newsId)
	if err != nil {
		_ = writeJsonError(w, http.StatusBadRequest, "Failed to parse ID")
		return
	}

	storage := StorageClient{}
	err = storage.Connect(viper.GetString("amqp_url"))
	if err != nil {
		_ = writeJsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	defer storage.Close()

	response, err := storage.GetNews(id)
	if err != nil {
		_ = writeJsonError(w, http.StatusInternalServerError, "Failed to get News")
		return
	}

	out, err := json.Marshal(response)
	_, _ = w.Write(out)
}

func AddNews(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	date := r.FormValue("date")

	_, err := time.Parse(time.RFC822Z, date)
	if err != nil {
		_ = writeJsonError(w, http.StatusBadRequest, "Failed to parse date")
		return
	}

	storage := StorageClient{}
	err = storage.Connect(viper.GetString("amqp_url"))
	if err != nil {
		_ = writeJsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	defer storage.Close()

	response, err := storage.AddNews(title, date)
	if err != nil {
		_ = writeJsonError(w, http.StatusInternalServerError, "Failed to add News")
		return
	}

	out, err := json.Marshal(response)
	_, _ = w.Write(out)
}

func writeJsonResponse(w http.ResponseWriter, status int, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		log.Println("Failed to marshal response")
		return err
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)

	return nil
}

func writeJsonError(w http.ResponseWriter, status int, msg string) error {
	log.Printf("Error: %s", msg)
	return writeJsonResponse(w, status, ErrorResponse{Code: status, Message: msg})
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
