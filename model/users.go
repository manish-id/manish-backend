package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID		  		primitive.ObjectID 			 `bson:"user_id,omitempty" json:"user_id"`
	Username 			string                		 `bson:"username" json:"username"`
	Password			string                		 `bson:"password" json:"password"`
	Email 				string                		 `bson:"email" json:"email"`
	RoleID 				string                		 `bson:"role" json:"role"` // "admin" or "user"
	CreatedAt 			primitive.DateTime       	 `bson:"created_at" json:"created_at"`
	UpdatedAt 			primitive.DateTime       	 `bson:"updated_at" json:"updated_at"`
}