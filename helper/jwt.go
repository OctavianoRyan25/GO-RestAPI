package helper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "rahasia"

func GenerateToken(id uint,email string) string {

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"id"	: id,
			"email"	: email,
		})

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error)  {
	errResponse := errors.New("sign into proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token ,_ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid{
		return nil, errResponse
	}
	return token.Claims.(jwt.MapClaims), nil
}
