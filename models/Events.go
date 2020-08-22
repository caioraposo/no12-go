package models

import (
    "github.com/Kamva/mgm/v2"
)

type Event struct {
    mgm.DefaultModel `bson:",inline"`
    Title            string `json:"title" bson:"title"`
    Description      string `json:"description" bson:"description"`
    Date             string `json:"date" bson:"date"`
}

func NewEvent(title, description string, date string) *Event {
    return &Event {
        Title:          title,
        Description:    description,
        Date:           date,
    }
}
