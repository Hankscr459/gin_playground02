package implement

import (
	"errors"
	"ginValid/dto/user"
	"ginValid/service"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserController struct {
	UserService service.UserService
}

var Email string

var IDUser string

func GenerJWT(t user.Read) (string, error) {
	secret := []byte(os.Getenv("JwtSecretKey"))

	payload := jwt.MapClaims{
		"email": t.Email,
		"name":  t.Name,
		"_id":   t.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func ProccessToken(tk string) (*user.Cliam, bool, string, error) {
	var u UserController
	secret := []byte(os.Getenv("JwtSecretKey"))
	claims := &user.Cliam{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Format token invalid")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err == nil {
		find, _ := u.UserService.GetUser(&claims.Name)
		if find != nil {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, true, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalid")
	}
	return claims, false, string(""), err
}
