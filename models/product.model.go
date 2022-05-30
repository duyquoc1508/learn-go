package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID  				primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name 				string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Category 		int 							 `json:"category,omitempty" bson:"category,omitempty"`
	Price 			float64            `json:"price,omitempty" bson:"pricate,omitempty"`

}
