package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const localSigningKey = "Local"
const TokenExpiredTime = 120 //Minute
type CustomClaims struct {
	UID       int
	Privilege int
	jwt.StandardClaims
}
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error = errors.New("TokenExpired")
	TokenInvalid     error = errors.New("TokenInvalid")
	TokenMalformed   error = errors.New("TokenMalformed")
	TokenNotValidYet error = errors.New("TokenNotValidYet")
)

func getSigningKey() string {
	return localSigningKey
}
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(getSigningKey()))
}
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = time.Now
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(TokenExpiredTime * time.Minute).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
func (j *JWT) GenerateToken(uid, privilege int) (token string, err error) {
	claims := CustomClaims{
		UID:       uid,
		Privilege: privilege,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(TokenExpiredTime * time.Minute).Unix()),
			Issuer:    getSigningKey(),
			NotBefore: int64(time.Now().Unix()),
		},
	}
	token, err = j.CreateToken(claims)
	return
}
