package main

import (
	"go_trial/gorest/chapter7/base62Example/base62"
	"log"
)

func main() {
	x := 100
	base62String := base62.Tobase62(x)
	log.Println(base62String)
	normalNumber := base62.Tobase10(base62String)
	log.Println(normalNumber)
}
