package repoImpl

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	config "go-demo/configs"
	constant "go-demo/constants"
	driver "go-demo/drivers"
	model "go-demo/models"
	repo "go-demo/repositories"
)

// Để 1 struct implement 1 interface thì struct đó phải implement toàn bộ method của interface đó đề cập
// Để khai báo struct UserRepoImpl implement interface IUserRepo thì struct UserRepoImpl phải implement toàn bộ các method của interface
type UserRepoImpl struct {
	Db *mongo.Database
}

// hàm này khai báo dữ liệu trả về là 1 interface. bên trong nó trả về 1 struct đã được implement
func NewUserRepo() repo.IUserRepo {
	return &UserRepoImpl{
		Db: driver.Mongo.Client.Database(config.DATABASE_NAME),
	}
}

func (mongo *UserRepoImpl) FindUserByEmail(email string) (model.User, error) {
	user := model.User{}
	result := mongo.Db.Collection(constant.USERS).FindOne(context.Background(), bson.M{"email": email})
	err := result.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (mongo *UserRepoImpl) CheckLoginInfo(email, password string) (model.User, error) {
	user := model.User{}
	result := mongo.Db.Collection(constant.USERS).FindOne(context.Background(), bson.M{"email": email, "password": password})
	err := result.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (mongo *UserRepoImpl) Insert(user *model.User) error {
	// muốn insert 1 struct vào mongodb thì phải sử dụng bson
	// bbytes, _ := user.MarshalBSON()
	_, err := mongo.Db.Collection(constant.USERS).InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
