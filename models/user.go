package models

type User struct {
	Name      string `binding:"required" label:"名稱" json:"name" bson:"name"`
	FirstName string `binding:"required" label:"名稱1" json:"first_name" bson:"first_name"`
	LastName  string `binding:"required" label:"名稱2" json:"last_name" bson:"last_name"`
	Age       uint8  `binding:"gte=0,lte=130" label:"年齡" json:"age" bson:"age"`
	Email     string `binding:"required" label:"電子郵件" json:"email" bson:"email"`
	Password  string `binding:"required" label:"密碼" json:"password" bson:"password"`
}
