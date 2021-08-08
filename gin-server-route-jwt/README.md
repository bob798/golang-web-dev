## gin-jwt 使用 demo

### 是什么
jwt 全称 JSON Web Token，用于网络中紧凑、安全的传输两端信息。
规范:RFC7519 https://datatracker.ietf.org/doc/html/rfc7519

jwt 解析工具https://jwt.io/
### 为什么要使用它
承载用户登录信息
### jwt vs session
session 存储在内容重
jwt 存储在客户端

jwt更利于处理扩展性、多端登录问题
参考：https://sherryhsu.medium.com/session-vs-token-based-authentication-11a6c5ac45e4

### 如何使用
#### 下载依赖包
```shell
go get github.com/appleboy/gin-jwt
```

#### 创建Token

````go
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

````

#### 解析Token

```go
// 解析 jwt
	parseToken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return secret, err
	})

	fmt.Printf("parseToken= %s\n", parseToken)
	// parseToken= &{eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4ODAzOTg2NzksImp0aSI6IjEiLCJpc3MiOiJib2IiLCJuYmYiOjE2MjgzOTg2NzksInN1YiI6InRlc3QifQ.AC5oLznuzMsWbOuclEfZwY2H6N0ExhO8lojXVwa7UAg %!s(*jwt.SigningMethodHMAC=&{HS256 5}) map[alg:HS256 typ:JWT] map[exp:%!s(float64=1.880398679e+09) iss:bob jti:1 nbf:%!s(float64=1.628398679e+09) sub:test] AC5oLznuzMsWbOuclEfZwY2H6N0ExhO8lojXVwa7UAg %!s(bool=false)}
	fmt.Printf("parseToken claim= %s\n", parseToken.Claims)
	// parseToken claim= map[exp:%!s(float64=1.880398679e+09) iss:bob jti:1 nbf:%!s(float64=1.628398679e+09) sub:test]

```