package util

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id       int32  `json:"user_id"`
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

func ValidarId(ctx *gin.Context, id int32) error {

	tokenHeader := ctx.GetHeader("authorization")
	fields := strings.Fields(tokenHeader)
	tokenValidate := fields[1]
	id_user := id
	textEncoded := strings.SplitAfter(tokenHeader[7:], ".")
	rawDecodedText, _ := base64.StdEncoding.DecodeString(textEncoded[1])
	rawDecodedTextConv := string(rawDecodedText)
	userIdConv := "userid" + strconv.Itoa(int(id_user))
	final := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(rawDecodedTextConv, "")
	match, _ := regexp.MatchString(userIdConv, final)
	errValidateId := ValidateToken(tokenValidate, ctx)

	if match != true {
		return fmt.Errorf("Voce não está autorizado a visualiuzar esses dados ")
	}

	return errValidateId
}

func ValidarUserName(ctx *gin.Context, name string) error {

	tokenHeader := ctx.GetHeader("authorization")
	fields := strings.Fields(tokenHeader)
	tokenValidate := fields[1]
	username := name
	textEncoded := strings.SplitAfter(tokenHeader[7:], ".")
	rawDecodedText, _ := base64.StdEncoding.DecodeString(textEncoded[1])
	rawDecodedTextConv := string(rawDecodedText)
	final := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(rawDecodedTextConv, "")
	match, _ := regexp.MatchString(username, final)
	errValidateUserName := ValidateToken(tokenValidate, ctx)

	if match != true {
		return fmt.Errorf("Voce não está autorizado a visualiuzar esses dados ")
	}

	return errValidateUserName
}
