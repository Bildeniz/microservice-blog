package main

import (
	"UserConsumer/db"
	"UserConsumer/models"
	"UserConsumer/repositories"
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s, %s", msg, err)
	}
}

func main() {
	ctx := context.Background()

	// Connect to DB
	db_conn := db.Connect()
	db.Migrate(db_conn, ctx)
	user_repo := repositories.NewUserRepository(db_conn)

	// Rabbitmq Connection
	conn, err := amqp.Dial("amqp://rabbitmquser:rabbitmqpass@rabbitmq:5672/")
	failOnError(err, "Failed to conntect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to oopen")
	defer ch.Close() // Close the connection after running function

	// Declaring the queue
	q, err := ch.QueueDeclare(
		"user-creat",
		true,  // durability
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)
	failOnError(err, "Failed to declare a queue")

	// Register to consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var user models.User

			err := json.Unmarshal([]byte(d.Body), &user)
			if err != nil {
				log.Print(err, "Failed to converting data to json")
				return
			}

			err = user_repo.CreateNewUser(&user)
			if err != nil {
				log.Print(err, "Failed to creating user on DB")
				return
			}

			log.Printf("%s %s %d User Created with succesfully", user.Name, user.Surname, user.ID)
			d.Ack(true)
		}
	}()

	log.Print("[*] Waiting for messages... To exit press CTRL+C")
	<-forever
}
