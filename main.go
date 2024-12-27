package main

import (
	"log"

	"github.com/gofiber/fiber"
)

type MongoInstance struct {
	Client string
	Db     string
}

var mg MongoInstance

const dbName = "fiber-hrms"

const mongoURI = "mongodb+srv://siddharth-tricon:Password@cluster0.ccgmbiv.mongodb.net/" + dbName

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {

}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {

	})

	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
