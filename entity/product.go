package entity

import "time"

type Product struct {
	ID          string    `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	Cost        int       `json:"cost" bson:"cost"`
	Quantity    int       `json:"quantity" bson:"quantity"`
	DateCreated time.Time `json:"date_created" bson:"dateCreated"`
	DateUpdated time.Time `json:"date_updated" bson:"dateUpdated"`
}
