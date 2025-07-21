/*
 * @Author: JimZhang
 * @Date: 2025-07-19 12:40:52
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-19 13:25:39
 * @FilePath: /server/utils/jwt.go
 * @Description:
 *
 */
package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var stSignKey = []byte(viper.GetString("jwt.signKey"))

type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, username string) (string, error) {

	iJwtCustClaims := &JwtCustClaims{
		ID:   id,
		Name: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Subject:  "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)

	return token.SignedString(stSignKey)
}

func ParseToken(tokenString string) (JwtCustClaims, error) {

	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("token is invalid")
	}
	return iJwtCustClaims, err
}

func IsVaildToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
