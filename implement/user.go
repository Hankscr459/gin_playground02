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

func (u *UserServiceImpl) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	create, err := u.usercollection.InsertOne(u.ctx, user)
	return create, err
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

func (u *UserServiceImpl) GetUserByEmail(email *string) (*user.Read, error) {
	var user *user.Read
	query := bson.D{bson.E{Key: "email", Value: email}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) UpdateUser(update *user.Update, userId *string) error {
	objID, idErr := primitive.ObjectIDFromHex(*userId)
	if idErr != nil {
		return idErr
	}
	query := bson.M{"_id": bson.M{"$eq": objID}}
	modify := bson.M{"$set": bson.M{"name": update.Name, "email": update.Email, "first_name": update.FirstName, "last_name": update.LastName, "age": update.Age}}
	res := u.usercollection.FindOneAndUpdate(u.ctx, query, modify)
	return res.Err()
}
