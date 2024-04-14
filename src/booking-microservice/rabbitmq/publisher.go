package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/common"
)

type RabbitMQ struct {
	Body      string
	QueueName string
}

func ConnectMQ() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(common.Config.RabbitMQ)
	if err != nil {
		log.Panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	return conn, ch
}

func CloseMQ(conn *amqp.Connection, ch *amqp.Channel) {
	conn.Close()
	ch.Close()
}

func (r *RabbitMQ) Publish(message interface{}) {

	conn, ch := ConnectMQ()
	defer CloseMQ(conn, ch)
	r.QueueName = common.Config.QueueName
	fmt.Println("QueueName: ", r.QueueName)
	q, err := ch.QueueDeclare(
		r.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	if err != nil {
		log.Panic(err)
	}
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		log.Panic(err)
	}
	log.Printf(" [x] Sent %s\n", body)
}
