package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/global"
	"go.uber.org/zap"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			zap.S().Info("请求未携带token，无权限访问")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请登录",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, TokenExpired) {
				zap.S().Info("token已过期")
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "token已过期",
				})
				c.Abort()
				return
			}
			zap.S().Info("token解析失败")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未登录",
			})
			c.Abort()
			return
		}
		c.Set("userId", claims.ID)
		c.Set("token", token)
		c.Next()
	}
}

type CustomClaims struct {
	ID int
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
	Expiration time.Duration
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.ServerConfig.JWT.SigningKey),
		Expiration: time.Duration(global.ServerConfig.JWT.Expiration) * time.Minute,
	}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(j.Expiration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
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
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		claims.StandardClaims.ExpiresAt = time.Now().Add(j.Expiration).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
