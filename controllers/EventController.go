package controllers

import (
    "os"
    "log"
    "fmt"

    "github.com/Kamva/mgm/v2"
    "github.com/gofiber/fiber"
    "github.com/caioraposo/no12-go/models"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

func init() {
    connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
    if len(connectionString) == 0 {
        connectionString = "mongodb://root:example@localhost:27017"
    }

    err := mgm.SetDefaultConfig(nil, "events", options.Client().ApplyURI(connectionString))
    if err != nil {
        log.Fatal(err)
    }
}

func GetAllEvents(ctx *fiber.Ctx) {
    collection := mgm.Coll(&models.Event{})
    events := []models.Event{}

    err := collection.SimpleFind(&events, bson.D{})
    if err != nil {
        ctx.Status(500).JSON(fiber.Map{
            "ok":    false,
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(fiber.Map{
        "ok":    true,
        "events": events,
    })
}
    

func CreateEvent(ctx *fiber.Ctx) {
    event := new(models.Event)

    if err := ctx.BodyParser(event); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("New event: %s\n", event.Title)
    if err := mgm.Coll(event).Create(event); err != nil {
        log.Fatal(err)
    }

}

func DeleteEvent(ctx *fiber.Ctx) {
    id := ctx.Params("id")
    event := &models.Event{}
    collection := mgm.Coll(event)

    _ = collection.FindByID(id, event)
    err := mgm.Coll(event).Delete(event)

    if err != nil {
        ctx.Status(404).JSON(fiber.Map{
            "ok": false,
            "error": "Event not found.",
        })
    }
}

func UpdateEvent(ctx *fiber.Ctx) {
    id := ctx.Params("id")
    event := &models.Event{}
    collection := mgm.Coll(event)

    err := collection.FindByID(id, event)
    if err != nil {
        ctx.Status(404).JSON(fiber.Map{
            "ok": false,
            "error": "Event not found.",
        })
    }

    if err := ctx.BodyParser(event); err != nil {
        log.Fatal(err)
    }
    _ = mgm.Coll(event).Update(event)
}

func GetEventByID(ctx *fiber.Ctx) {
    id := ctx.Params("id")

    event := &models.Event{}
    collection := mgm.Coll(event)

    err := collection.FindByID(id, event)
    if err != nil {
        ctx.Status(404).JSON(fiber.Map{
            "ok": false,
            "error": "Event not found.",
        })
    }

    ctx.JSON(fiber.Map{
        "ok": true,
        "event": event,
    })
}
