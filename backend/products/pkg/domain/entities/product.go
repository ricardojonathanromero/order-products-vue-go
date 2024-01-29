package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Sku         string             `json:"sku" bson:"sku"`
	Price       float64            `json:"price" bson:"price"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Quantity    int64              `json:"quantity" bson:"quantity"`
	Image       string             `json:"image" bson:"image"`
	Category    string             `json:"categories" bson:"categories"`
}
