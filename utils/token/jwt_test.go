package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)


func TestJWT(t *testing.T) {

	claims := jwt.MapClaims{
		"exp" : time.Now().Add(time.Hour * 72).Unix(),
		"iat" : time.Now().Unix(),
		"data": "string(jsonStr)",
	}

	tokenstr,ok := NewToken(claims)
	if !ok {
		t.Fatal("token 生成错误")
	}

	data ,ok := CheckToken(tokenstr)
	if !ok {
		t.Fatal("token 验证错误")
	}
	t.Log("token check success",data)
}
