package main

import "github.com/streadway/amqp"

type Workers struct {
	conn amqp.Connection
}
