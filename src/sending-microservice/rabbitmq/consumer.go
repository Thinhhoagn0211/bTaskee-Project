package rabbitmq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/common"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/modules"
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

func (r *RabbitMQ) Consume() {

	conn, ch := ConnectMQ()
	defer CloseMQ(conn, ch)

	r.QueueName = common.Config.QueueName
	// log.Println(ch)
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Panic(err)
	}
	k := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("1.Received a message: %s", d.Body)
			// d.Ack(false)
			var newJob modules.Jobs

			err := json.Unmarshal(d.Body, &newJob.Job)
			if err != nil {
				log.Println("Error unmarshaling message body:", err)
				continue
			}
			newJob.InsertEmployeeToService(newJob)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-k
}
