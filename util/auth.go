package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken(token string, ctx *gin.Context) error {

	claims := &Claims{}
	var jwtSigned = []byte("secret_key")
	tokenParse, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSigned, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return err
		}
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}
	if !tokenParse.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token is Invalid")
		return nil
	}
	ctx.Next()
	return nil
}

func GetTokenInHeaderAndVerify(ctx *gin.Context) error {
	authorizationHearderKey := ctx.GetHeader("authorization")
	fields := strings.Fields(authorizationHearderKey)
	tokenValidate := fields[1]
	errValidateToken := ValidateToken(tokenValidate, ctx)
	if errValidateToken != nil {
		return errValidateToken
	}
	return nil
}
