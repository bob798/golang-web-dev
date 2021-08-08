package main

import (
	"fmt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"time"
)

func main() {

	// 创建 声明
	claim := jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Unix() + time.Hour.Microseconds()/100*7,
		Id:        "1",
		IssuedAt:  0,
		Issuer:    "bob",
		NotBefore: time.Now().Unix(),
		Subject:   "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 我是服务端私钥
	secret := []byte("secret")
	// 创建 jwt
	t, err := token.SignedString(secret)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token= %s\n", t)
	// token= eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4ODAzOTg2NzksImp0aSI6IjEiLCJpc3MiOiJib2IiLCJuYmYiOjE2MjgzOTg2NzksInN1YiI6InRlc3QifQ.AC5oLznuzMsWbOuclEfZwY2H6N0ExhO8lojXVwa7UAg

	// 解析 jwt
	parseToken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return secret, err
	})

	fmt.Printf("parseToken= %s\n", parseToken)
	// parseToken= &{eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4ODAzOTg2NzksImp0aSI6IjEiLCJpc3MiOiJib2IiLCJuYmYiOjE2MjgzOTg2NzksInN1YiI6InRlc3QifQ.AC5oLznuzMsWbOuclEfZwY2H6N0ExhO8lojXVwa7UAg %!s(*jwt.SigningMethodHMAC=&{HS256 5}) map[alg:HS256 typ:JWT] map[exp:%!s(float64=1.880398679e+09) iss:bob jti:1 nbf:%!s(float64=1.628398679e+09) sub:test] AC5oLznuzMsWbOuclEfZwY2H6N0ExhO8lojXVwa7UAg %!s(bool=false)}
	fmt.Printf("parseToken claim= %s\n", parseToken.Claims)
	// parseToken claim= map[exp:%!s(float64=1.880398679e+09) iss:bob jti:1 nbf:%!s(float64=1.628398679e+09) sub:test]
}
