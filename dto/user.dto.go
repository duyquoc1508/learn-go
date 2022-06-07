package dto // data transfer object

type Register struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
