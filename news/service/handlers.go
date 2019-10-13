package service

import (
	"crypto/sha1"
	"encoding/json"
	contract "github.com/Delisa-sama/go-test/proto"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
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

	request := &contract.GetNewsRequest{Id: uint32(id)}

	out, err := proto.Marshal(request)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to encode contract")
		return
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to connect to broker")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to get channel")
		return
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to declare queue")
		return
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to subscribe to queue")
		return
	}

	h := sha1.New()
	h.Write([]byte(string(time.Now().UnixNano())))
	correlationId := string(h.Sum(nil))

	err = ch.Publish(
		"",
		"load",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: correlationId,
			ReplyTo:       q.Name,
			Body:          out,
		})

	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to send RPC")
		return
	}

	listen := make(chan []byte)
	errorCh := make(chan string)

	go func() {
		for d := range msgs {
			if correlationId == d.CorrelationId {
				rpcResponse := &contract.OnGetNewsResponse{}
				if err := proto.Unmarshal(d.Body, rpcResponse); err != nil {
					log.Println("Failed to parse rpcResponse")
					continue
				}
				if rpcResponse.Status != contract.ResponseStatus_OK {
					errorCh <- "Failed to get News"
				}
				log.Printf("Received rpcResponse on save: %s", rpcResponse)
				response := getNewsResponse{Id: rpcResponse.Id, Title: rpcResponse.Title, Date: rpcResponse.Date}
				data, _ := json.Marshal(response)
				listen <- data
				break
			}
		}
	}()

	select {
	case data := <-listen:
		_, _ = w.Write(data)
	case <-time.After(5 * time.Second): // TODO: Timeout from config
		writeJsonError(w, http.StatusGatewayTimeout, "Timeout")
	case msg := <-errorCh:
		writeJsonError(w, http.StatusInternalServerError, msg)
	}
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

	request := &contract.AddNewsRequest{
		Title: title,
		Date:  date,
	}

	out, err := proto.Marshal(request)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to encode contract")
		return
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to connect to broker")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to get channel")
		return
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to declare queue")
		return
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to subscribe to queue")
		return
	}

	h := sha1.New()
	h.Write([]byte(string(time.Now().UnixNano())))
	correlationId := string(h.Sum(nil))

	err = ch.Publish(
		"",
		"save",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: correlationId,
			ReplyTo:       q.Name,
			Body:          out,
		})

	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Failed to send RPC")
		return
	}

	listen := make(chan []byte)
	errorCh := make(chan string)

	go func() {
		for d := range msgs {
			if correlationId == d.CorrelationId {
				rpcResponse := &contract.OnAddNewsResponse{}
				if err := proto.Unmarshal(d.Body, rpcResponse); err != nil {
					log.Fatalln("Failed to parse rpc_response")
				}
				if rpcResponse.Status != contract.ResponseStatus_OK {
					errorCh <- "Failed to add News"
				}
				log.Printf("Received rpc_response on save: %s", rpcResponse)

				response := &addNewsResponse{Id: rpcResponse.Id}

				data, _ := json.Marshal(response)
				listen <- data
				break
			}
		}
	}()

	select {
	case data := <-listen:
		_, _ = w.Write(data)
	case <-time.After(5 * time.Second):
		writeJsonError(w, http.StatusGatewayTimeout, "Timeout")
	case msg := <-errorCh:
		writeJsonError(w, http.StatusInternalServerError, msg)
	}
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
	writeJsonResponse(w, status, errorResponse{Code: status, Message: msg})
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type addNewsResponse struct {
	Id uint32 `json:"id"`
}

type getNewsResponse struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}
