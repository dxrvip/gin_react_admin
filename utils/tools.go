package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

// 工具包

type MyCustomClaims struct {
	Foo string `json:"foo"`
	Id  int    `json:"id"`
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(username string, userId int) (string, error) {
	// 生成token
	claims := MyCustomClaims{
		username,
		userId,
		jwt.RegisteredClaims{
			// 设置签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			// NumericDate 也可以使用固定日期 设置过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signKey := viper.GetString("app.Key")

	ss, err := token.SignedString([]byte(signKey))
	fmt.Printf("token: %s \n", ss)
	return ss, err
}

// 校验jwt
func VerifyJWT(tokenString string, signingKey []byte) (*MyCustomClaims, error) {
	// 解析JWT
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 从 token 中获取自定义 claims
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		// 检查是否过期
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			return nil, fmt.Errorf("JWT token has expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid JWT token")
}

// 加密密码
func EncryptPassword(password string) string {
	//哈希算法
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// 密码校验
func CheckPassword(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 生成随机字符串
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	result := make([]byte, length)
	for i := range result {

		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
