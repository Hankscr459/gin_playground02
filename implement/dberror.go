package implement

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Error struct {
	Message string `json:"message"`
}

func ValidDbError(err error, ctx *gin.Context) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		out := make([]Error, 1)
		for _, we := range e.WriteErrors {
			if we.Message != "" && we.Code == 11000 {

				if strings.Index(we.Message, "name_1") != -1 {
					out[0] = Error{"名稱已被註冊過"}
					ctx.JSON(http.StatusBadGateway, gin.H{"errors": out})
				} else if strings.Index(we.Message, "email_1") != -1 {
					out[0] = Error{"Email已被註冊過"}
					ctx.JSON(http.StatusBadGateway, gin.H{"errors": out})
				} else {
					out[0] = Error{we.Message}
					ctx.JSON(http.StatusBadGateway, gin.H{"errors": out})
				}
			} else {
				out[0] = Error{we.Message}
				ctx.JSON(http.StatusBadGateway, gin.H{"errors": out})
			}
		}
		return true
	} else {
		return false
	}
}
