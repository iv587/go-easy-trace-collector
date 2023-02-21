package auth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secret     = "apmlog"
	expireTime = 3600 * 2
)

type jwtClaims struct {
	jwt.StandardClaims
	UserName string
	Password string
}

// Token 生成token
func Token(name, password string) (string, error) {
	claims := new(jwtClaims)
	claims.UserName = name
	claims.Password = password
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("网络错误，请重试")
	}
	return signedToken, nil
}

// Verify 校验token
func Verify(strToken string) (bool, error) {
	if strToken == "" {
		return false, nil
	}
	token, err := jwt.ParseWithClaims(strToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, errors.New("网络错误，请重试")
	}

	_, ok := token.Claims.(*jwtClaims)
	if !ok {
		return false, nil
	}
	if err := token.Claims.Valid(); err != nil {
		return false, nil
	}
	return true, nil
}
