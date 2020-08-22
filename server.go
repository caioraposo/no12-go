package main

import "github.com/gofiber/fiber"
import "github.com/caioraposo/no12-go/controllers"

func main() {
    app := fiber.New()

    app.Get("/api/events", controllers.GetAllEvents)
    app.Get("/api/events/:id", controllers.GetEventByID)
    app.Post("/api/events", controllers.CreateEvent)
    app.Patch("/api/events/:id", controllers.UpdateEvent)
    app.Delete("/api/events/:id", controllers.DeleteEvent)

    app.Listen(3000)
}

