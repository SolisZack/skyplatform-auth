package method

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"skyplatform-auth/global"
	"skyplatform-auth/request"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.JWT.SigningKey),
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors & jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors & jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

func (j *JWT) GetETCDJWT(userName string) (string, error){
	resp, err := global.ETCD.Get(context.Background(), global.UserPrefix+ userName)
	if err != nil{
		return "", err
	}else if len(resp.Kvs) <= 0{
		return "", ErrorKeyNotFound
	}

	log.Printf("Get UserJWT, Key: %s : Value: %s\n", resp.Kvs[0].Key, resp.Kvs[0].Value)
	return string(resp.Kvs[0].Value), nil
}

func (j *JWT) SetETCDJWT(userName string, token string) error {
	_, err := global.ETCD.Put(context.Background(), global.UserPrefix+ userName, token)
	if err != nil{
		return err
	}

	return nil
}