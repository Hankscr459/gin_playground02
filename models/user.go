package user

type User struct {
	FirstName string `binding:"required" label:"名稱1" json:"first_name"`
	LastName  string `binding:"required" label:"名稱2" json:"last_name"`
	Age       uint8  `binding:"gte=0,lte=130" label:"年齡" json:"age"`
	Email     string `binding:"required" label:"電子郵件" json:"email"`
}
