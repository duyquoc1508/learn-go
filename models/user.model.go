package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` // tag golang
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"-" bson:"password"`
	DisplayName string             `json:"displayName" bson:"displayName"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// Tag golang (json, bson) có tác dụng mapping struct User vào trong kiểu dữ liệu json hoặc bson. Mongodb là bson
// set the Password field to never Marshal into JSON. `json:"-"`

// implement auto create `createdAt` add `updatedAt` because mongoDB driver not support this
// ref: https://stackoverflow.com/questions/71902455/autofill-created-at-and-updated-at-in-golang-struct-while-pushing-into-mongodb
// MarshalBSON() function will be called when you save values of your `*User` type
// IMPORTANCE: Note the method has pointer receiver (*User), so use a pointer to your value when save
func (user *User) MarshalBSON() ([]byte, error) {
	// fmt.Println("marshal bson")
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	user.UpdatedAt = time.Now()

	type my User
	return bson.Marshal((*my)(user))
}
