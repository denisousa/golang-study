package main

import (
	"net/http"

	"github.com/gofiber/fiber"
)

type Student struct {
	Name string
	Age  uint8
	ClassName string
	NoteGeral uint8
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("This is the GET")
		// http.StatusNotFound
	})

	app.Get("/student/:name", func(ctx *fiber.Ctx) {
		ctx.Send("This is the GET")
		name := ctx.Params("name")
		ctx.Send("The name of student: " + name)
		// http.StatusNotFound
	})

	app.Post("/", func(ctx *fiber.Ctx) {
		ctx.Status(http.StatusCreated).Send("This is the POST")
	})

	app.Put("/", func(ctx *fiber.Ctx) {
		ctx.Send("Olá, Mundo!")
	})

	app.Delete("/", func(ctx *fiber.Ctx) {
		ctx.Send("Olá, Mundo!")
	})

	app.Patch("/", func(ctx *fiber.Ctx) {
		ctx.Send("Olá, Mundo!")
	})

	app.Listen(8080)
}
