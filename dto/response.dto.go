package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty -> nếu field này có giá trị là empty thì bỏ nó để nó không xuất hiện trong json
}
