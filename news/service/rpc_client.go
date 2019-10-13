package service

import (
	"crypto/sha1"
	"errors"
	contract "github.com/Delisa-sama/go-test/proto"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type StorageClient struct {
	conn *amqp.Connection
}

func (s *StorageClient) Connect(connectionString string) error {
	var err error
	s.conn, err = amqp.Dial(connectionString)
	if err != nil {
		log.Println("Failed to connect to broker")
		return err
	}

	return nil
}

func (s *StorageClient) Close() {
	_ = s.conn.Close()
}

func (s *StorageClient) GetNews(id int) (*GetNewsResponse, error) {
	request := &contract.GetNewsRequest{Id: uint32(id)}
	out, err := proto.Marshal(request)
	if err != nil {
		log.Println("Failed to encode contract")
		return nil, err
	}

	if s.conn == nil {
		log.Println("Connection is nil")
		return nil, errors.New("connection is nil")
	}

	ch, err := s.conn.Channel()
	if err != nil {
		log.Println("Failed to get channel")
		return nil, err
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Println("Failed to declare queue")
		return nil, err
	}

	// Consume queue for receive Storage-Service reply
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println("Failed to subscribe to queue")
		return nil, err
	}

	// Generate correlation ID
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
		log.Println("Failed to send RPC")
		return nil, err
	}

	listen := make(chan *GetNewsResponse) // channel for good response
	errorCh := make(chan string)          // channel for error response

	go func() {
		for d := range msgs {
			if correlationId == d.CorrelationId {
				rpcResponse := &contract.OnGetNewsResponse{}
				if err := proto.Unmarshal(d.Body, rpcResponse); err != nil {
					log.Println("Failed to parse rpcResponse")
					continue
				}
				if rpcResponse.Status != contract.ResponseStatus_OK {
					errorCh <- "failed to get News"
				}
				log.Printf("Received rpcResponse on save: %s", rpcResponse)
				response := &GetNewsResponse{Id: rpcResponse.Id, Title: rpcResponse.Title, Date: rpcResponse.Date}
				listen <- response
				break
			}
		}
	}()

	select {
	case response := <-listen:
		return response, nil
	case <-time.After(time.Duration(viper.GetInt("storage_timeout")) * time.Second):
		return nil, errors.New("timeout")
	case msg := <-errorCh:
		return nil, errors.New(msg)
	}
}

func (s *StorageClient) AddNews(title, date string) (*AddNewsResponse, error) {
	// Prepare RPC Request
	rpcRequest := &contract.AddNewsRequest{
		Title: title,
		Date:  date,
	}

	out, err := proto.Marshal(rpcRequest)
	if err != nil {
		log.Println("Failed to encode contract")
		return nil, err
	}

	if s.conn == nil {
		log.Println("Connection is nil")
		return nil, errors.New("connection is nil")
	}

	ch, err := s.conn.Channel()
	if err != nil {
		log.Println("Failed to get channel")
		return nil, err
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Println("Failed to declare queue")
		return nil, err
	}

	// Consume queue for receive Storage-Service reply
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println("Failed to subscribe to queue")
		return nil, err
	}

	// Generate correlation ID
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
		log.Println("Failed to send RPC")
		return nil, err
	}

	listen := make(chan *AddNewsResponse) // channel for good response
	errorCh := make(chan string)          // channel for error response

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

				response := &AddNewsResponse{Id: rpcResponse.Id}
				listen <- response
				break
			}
		}
	}()

	select {
	case data := <-listen:
		return data, nil
	case <-time.After(time.Duration(viper.GetInt("storage_timeout")) * time.Second): // TODO: Get timeout from CLI
		return nil, errors.New("timeout")
	case msg := <-errorCh:
		return nil, errors.New(msg)
	}
}
