package valid

import (
	"errors"
	"net/http"
	"reflect"

	info "ginValid/dto/info"
	test "ginValid/dto/test"
	models "ginValid/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ValidError(err error, c *gin.Context, model interface{}) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]info.Error, len(ve))
		for i, fe := range ve {
			Error := info.Error{}
			var errorMessage string = ""
			field, ok := reflect.TypeOf(model).Elem().FieldByName(fe.Field())
			if !ok {
				panic("Field not found")
			}
			label := string(field.Tag.Get("label"))
			switch fe.Tag() {
			case "required":
				errorMessage = label + " 欄位必填"
			case "email":
				errorMessage = "Email 格式不對"
			default:
				errorMessage = fe.Error()
			}
			Error.Message = errorMessage
			out[i] = Error
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		c.Abort()
	}
}

func SignupValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
			ValidError(err, c, &user)
		} else {
			c.Set("user", user)
			c.Next()
		}
	}
}

func DateValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var date test.DateBetween
		if err := c.ShouldBindQuery(&date); err == nil {
			c.Set("date", date)
			c.Next()
		} else {
			value, isErr := err.(validator.ValidationErrors)
			if isErr {
				ValidError(value, c, &date)
			}
		}
	}
}
