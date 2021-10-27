package main

import (
	"fmt"

	"github.com/klintonlee/rabbitmq-poc/internal/rabbitmq"
)

// App - the struct which contains things like pointer
// to our internal RabbitMQ Service
type App struct {
	Rmq *rabbitmq.RabbitMQ
}

// Run - handles the startup of our application
func Run() error {
	fmt.Println("Go RabbitMQ crash course")

	rmq := rabbitmq.NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	if err := app.Rmq.Connect(); err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()

	if err := app.Rmq.Publish("Hello World"); err != nil {
		fmt.Println(err)
		return err
	}

	app.Rmq.Consume()

	return nil
}

// main - the entrypoint for the application
func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up our application")
		fmt.Println(err)
	}
}
