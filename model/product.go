package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ProductID		  		primitive.ObjectID 			`bson:"product_id,omitempty" json:"product_id"`
	CategoryID 				primitive.ObjectID 			`bson:"category_id" json:"category_id"`
	ProductName 			string                	 	`bson:"product_name" json:"product_name"`
	ProductPrice 			float64                 	`bson:"product_price" json:"product_price"`
	Stock 					int                   		`bson:"stock" json:"stock"`
	ProductImage 			string                	 	`bson:"product_image" json:"product_image"`
	Size 					string                	 	`bson:"size" json:"size"`
	Details 				string                	 	`bson:"details" json:"details"`
	CreatedAt				primitive.DateTime			`bson:"created_at" json:"created_at"`
}