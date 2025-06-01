package api

import (
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims defines the structure of JWT claims used in the application.
type Claims struct {
	// UserId   int    `json:"user_id"`
	// Username string `json:"user_name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// JwtMiddleware handles JWT authentication using a shared secret.
type JwtMiddleware struct {
	secret []byte
}

// NewJwtMiddleware creates a new instance of JwtMiddleware with the provided secret.
func NewJwtMiddleware(secret string) *JwtMiddleware {
	return &JwtMiddleware{
		secret: []byte(secret),
	}
}

// Auth returns a Gin middleware function that validates JWT tokens from the Authorization header.
func (j *JwtMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responser.ResponseError(c, bizcode.UNAUTHORIZED)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			responser.ResponseError(c, bizcode.UNAUTHORIZED)
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return j.secret, nil
		})

		if err != nil || !token.Valid {
			responser.ResponseError(c, bizcode.UNAUTHORIZED)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			responser.ResponseError(c, bizcode.UNAUTHORIZED)
			c.Abort()
			return
		}

		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			responser.ResponseError(c, bizcode.UNAUTHORIZED)
			c.Abort()
			return
		}

		// c.Set("user_id", claims.UserID)
		// c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("claims", claims)

		c.Next()
	}
}

// GetUserID retrieves the user ID from the Gin context.
// func GetUserID(c *gin.Context) (string, bool) {
// 	userID, exists := c.Get("user_id")
// 	if !exists {
// 		return "", false
// 	}
// 	return userID.(string), true
// }

// GetUsername retrieves the username from the Gin context.
// func GetUsername(c *gin.Context) (string, bool) {
// 	username, exists := c.Get("username")
// 	if !exists {
// 		return "", false
// 	}
// 	return username.(string), true
// }

// GetRole retrieves the role from the Gin context.
func GetRole(c *gin.Context) (string, bool) {
	role, exists := c.Get("role")
	if !exists {
		return "", false
	}
	return role.(string), true
}

// GetClaims retrieves the full JWT claims from the Gin context.
func GetClaims(c *gin.Context) (*Claims, bool) {
	claims, exists := c.Get("claims")
	if !exists {
		return nil, false
	}
	return claims.(*Claims), true
}
