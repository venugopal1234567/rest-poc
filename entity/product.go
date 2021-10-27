package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SkuId    string             `json:"skuId" bson:"skuId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Cost     int                `json:"cost" bson:"cost,omitempty"`
	Quantity int                `json:"quantity" bson:"quantity,omitempty"`
}
