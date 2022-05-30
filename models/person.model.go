package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person represents a person document in MongoDB
/* Giải thích phần struct tags. Phần phía sau gọi là struct tags
- `json` là thể hiện trong code, ví dụ dữ liệu đầu vào đầu ra sẽ là object có key giống như phần json
- `bson` là thể hiện trong database. Tên cột trong database sẽ giống phần bson

Nếu không có mapping thì golang sẽ tự hiểu như sau:
The Go Driver generates BSON key using the lowercase of the corresponding struct field
- `json` là key của struct luôn. Dữ liệu đầu ra sẽ là object có key giống như key của struct. Giống y chang định dạng hoa thường
- `bson` là lowercase key của struct. Data ở db sẽ là object có property là lowercase từng key của struct

Tóm lại: `bson` thể hiện dữ liệu trong database, `json` thể hiện dữ liệu trong code
*/
// Quá trình chuyển đổi giữa giá trị Go <-> BSON.
// Go -> BSON: được gọi là marshalling
// BSON -> Go: được gọi là unmarshalling
type Person struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Age         int                `json:"age,omitempty" bson:"age,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	// PriceOfItem       int  => nếu không có mapping thì golang sẽ ngầm hiểu -> `json:"PriceOfItem",omitempty bson: "priceofitem",omitempty`
}

// ref: https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/
// + omitempty:
