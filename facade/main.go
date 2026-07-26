package main

import (
	"facade/phone_number"
)

func main() {
	consumer := phone_number.NewConsumer()
	consumer.Process("0044445635")
}
