package extension

import (
	"context"
	"fmt"
	c "ginValid/controller"
	"ginValid/implement"
	"ginValid/service"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	userservice    service.UserService
	usercontroller c.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func ConnectDb() c.UserController {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI(os.Getenv("MongoApplyURI"))
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection establish")
	usercollection = mongoclient.Database("userdb").Collection("users")
	for _, f := range []string{"name", "email"} {
		_, err := usercollection.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: f, Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	userservice = implement.NewUserService(usercollection, ctx)
	usercontroller = c.New(userservice)
	return usercontroller
}

func Disconnect() {
	mongoclient.Disconnect(ctx)
}
