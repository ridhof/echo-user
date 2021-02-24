package middlewares

import (
	"echo-user/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// CreateToken to create new token session by combining userId with other data.
func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

// ExtractTokenUserID to extract User Id from given token string.
func ExtractTokenUserID(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["userID"].(int)
		return userID
	}
	return 0
}
