package implement

import (
	"context"
	"ginValid/dto/user"
	"ginValid/models"
	"ginValid/service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) service.UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*user.Read, error) {
	var user *user.Read
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetUserById(id *string) (*user.Read, error) {
	var user *user.Read
	objID, idErr := primitive.ObjectIDFromHex(*id)
	if idErr != nil {
		return nil, idErr
	}
	query := bson.M{"_id": bson.M{"$eq": objID}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

// convert id string to ObjectId .FindId(bson.M{"_id": bson.ObjectIdHex("56bdd27ecfa93bfe3d35047d")})
// objectId, err := primitive.ObjectIDFromHex("5b9223c86486b341ea76910c")
// if err != nil{
//     log.Println("Invalid id")
// }
