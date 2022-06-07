package driver

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB() *MongoDB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	connectStr := os.Getenv("MONGODB_URI")
	if connectStr == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable")
	}
	// Set client option
	clientOptions := options.Client().ApplyURI(connectStr)

	// `ctx`: Limit 1 khoảng thời gian làm 1 việc gì đó. Nếu trong khoảng thời gian này không thực hiện tác vụ đó thành công thì nó sẽ hủy
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to database")

	Mongo.Client = client
	return Mongo
}
