package repository

import (
	// giống như import trong js. prefix là alias của package sử dụng trong file này. Nếu không có thì go sẽ mặc định sử dụng tên sau dấu / cuối cùng. Trường hợp này là 'model'
	model "go-demo/models"
)

type IUserRepo interface {
	FindUserByEmail(email string) (model.User, error)
	CheckLoginInfo(email string, password string) (model.User, error)
	Insert(user model.User) error
}
