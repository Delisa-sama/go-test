package main

import (
	"flag"
	contract "github.com/Delisa-sama/go-test/proto"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type News struct {
	gorm.Model
	Title string
	Date  int64
}

func initFlags() error {
	flag.String("db_dialect", "sqlite3", "Dialect of DB for gorm")
	flag.String("db_args", "./.db", "Args for DB. For sqlite it is path to db file, for PostgreSQL it is connection string")
	flag.String("amqp_url", "amqp://guest:guest@localhost:5672/", "URL string for connect to message broker via amqp")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	return viper.BindPFlags(pflag.CommandLine)
}

func main() {
	err := initFlags()
	if err != nil {
		log.Fatalln("Failed to parse flags")
	}

	db, err := gorm.Open(viper.GetString("db_dialect"), viper.GetString("db_args"))
	if err != nil {
		log.Fatalln("Failed to connect to DB")
	}
	defer db.Close()

	db.AutoMigrate(&News{})

	conn, err := amqp.Dial(viper.GetString("amqp_url"))
	if err != nil {
		log.Fatalln("Failed to connect to broker")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("Failed to open a channel")
	}
	defer ch.Close()

	qSave, err := ch.QueueDeclare(
		"save", // name
		false,  // durable
		false,  // delete when usused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalln("Failed to declare a queue")
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalln("Failed to set QoS")
	}

	saves, err := ch.Consume(
		qSave.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalln("Failed to register a consumer")
	}

	qLoad, err := ch.QueueDeclare(
		"load", // name
		false,  // durable
		false,  // delete when usused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalln("Failed to declare a queue")
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalln("Failed to set QoS")
	}

	loads, err := ch.Consume(
		qLoad.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalln("Failed to register a consumer")
	}

	listen := make(chan bool)

	go func() {
		for d := range saves {
			request := &contract.AddNewsRequest{}
			response := &contract.OnAddNewsResponse{}

			if err := proto.Unmarshal(d.Body, request); err != nil {
				log.Println("Failed to parse contract")
				response.Status = contract.ResponseStatus_FAIL
				err = PublishResponse(ch, &d, response)
				if err != nil {
					log.Println("Failed to publish response")
				}
				continue
			}

			log.Printf("Received a message: %s", request)

			date, err := time.Parse(time.RFC822Z, request.Date)
			if err != nil {
				log.Println("Failed to parse date")
				response.Status = contract.ResponseStatus_FAIL
				err = PublishResponse(ch, &d, response)
				if err != nil {
					log.Println("Failed to publish response")
				}
				continue
			}

			news := News{Title: request.Title, Date: int64(date.Unix())}
			db.NewRecord(news)
			db.Create(&news)

			response.Status = contract.ResponseStatus_OK
			response.Id = uint32(news.ID)

			err = PublishResponse(ch, &d, response)
			if err != nil {
				log.Println("Failed to publish response: ", err)
				continue
			}

			_ = d.Ack(false)
			log.Println("Replied: ", response)
		}
	}()

	go func() {
		for d := range loads {
			request := &contract.GetNewsRequest{}
			response := &contract.OnGetNewsResponse{}

			if err := proto.Unmarshal(d.Body, request); err != nil {
				log.Println("Failed to parse contract")
				response.Status = contract.ResponseStatus_FAIL
				err = PublishResponse(ch, &d, response)
				if err != nil {
					log.Println("Failed to publish response")
				}
				continue
			}

			log.Printf("Received request for load News: %s", request)

			news := News{}

			if err := db.First(&news, request.Id).Error; gorm.IsRecordNotFoundError(err) {
				log.Println("Record not found")
				response.Status = contract.ResponseStatus_FAIL
				err = PublishResponse(ch, &d, response)
				if err != nil {
					log.Println("Failed to publish response")
				}
				continue
			}

			date := time.Unix(news.Date, 0)

			response.Status = contract.ResponseStatus_OK
			response.Id = uint32(news.ID)
			response.Title = news.Title
			response.Date = date.Format(time.RFC822Z)

			err = PublishResponse(ch, &d, response)

			if err != nil {
				log.Println("Failed to publish response: ", err)
				continue
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-listen
}

func PublishResponse(ch *amqp.Channel, delivery *amqp.Delivery, response proto.Message) error {
	body, err := proto.Marshal(response)
	if err != nil {
		log.Println("Failed to parse response")
		return err
	}

	_ = delivery.Ack(false)

	err = ch.Publish(
		"",
		delivery.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: delivery.CorrelationId,
			Body:          body,
		},
	)

	if err != nil {
		log.Fatalln("Failed to publish response")
	}

	return nil
}
