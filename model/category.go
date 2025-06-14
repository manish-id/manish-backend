package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	CategoryID 				primitive.ObjectID 			`bson:"category_id" json:"category_id"`
	CategoryName 			string                	 	`bson:"category_name" json:"category_name"`
	CreatedAt				primitive.DateTime			`bson:"created_at" json:"created_at"`
	UpdatedAt				primitive.DateTime			`bson:"updated_at" json:"updated_at"`
}