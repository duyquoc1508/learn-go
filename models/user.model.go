package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"-" bson:"password"`
	DisplayName string             `json:"displayName" bson:"displayName"`
}

// Tag golang (json, bson) có tác dụng mapping struct User vào trong kiểu dữ liệu json hoặc bson. Mongodb là bson
// set the Password field to never Marshal into JSON. `json:"-"`
