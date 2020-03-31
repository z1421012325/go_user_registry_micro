package token

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

/*
	jwt.SigningMethodHS256 和 jwt.SigningMethodHMAC 一对
 */

var (
	SecretKey = ""
)

func NewSecretKey()[]byte{
	if len(os.Getenv("SecretKey")) == 0 {
		SecretKey = "xzc45sv1r3&2xd@sc&*14a~~~sdzRZXv"
	}
	return []byte(SecretKey)
}


func NewToken(v interface{}) (string,bool) {

	jsonStr,err := json.Marshal(v)
	if err != nil {
		return "",false
	}

	claims := jwt.MapClaims{
		"exp" : time.Now().Add(time.Hour * 72).Unix(),
		"iat" : time.Now().Unix(),
		"data": string(jsonStr),
	}

	//                                       算法            负载payload               盐值
	tokenString,err := jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString(NewSecretKey())
	if err != nil {
		return "",false
	}

	return tokenString,true
}






func CheckToken(tokenString string) (string,bool){
	// 检测算法,返回盐值解析得到token struct对象
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return NewSecretKey(), nil
	})
	if err != nil {
		return "", false
	}

	// 解析payload类型				MapClaims
	if claims ,ok := token.Claims.(jwt.MapClaims); ok && int64(claims["exp"].(float64)) >= time.Now().Unix() {
		return claims["data"].(string),true
	}
	return "",false
}

