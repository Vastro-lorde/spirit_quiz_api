package middlewares

import (
	"errors"
	"net/http"
	"spirit_quiz/config"
	dtos "spirit_quiz/internal/Dtos"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var configs, err = config.GetConfigs()

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err != nil {
			context.JSON(http.StatusUnauthorized, err)
			context.Abort()
		}

		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			context.Abort()
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token signing method")
			}
			return []byte(configs.JWTSecret), nil
		})

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			context.Abort()
			return
		}

		iss, ok := claims["iss"].(string)
		if !ok || iss != configs.JWTIssuer {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		aud, ok := claims["aud"].(string)
		if !ok || aud != configs.JWTAudience {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ID, ok := claims["ID"].(string)
		Role, ok := claims["Role"].(string)
		Email, ok := claims["Email"].(string)
		var user = dtos.UserClaim{
			ID:    ID,
			Email: Email,
			Role:  Role,
		}
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}

		context.Set("user", &user)
		context.Next()
	}
}

func RoleAuth(role string) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, exists := context.Get("user")
		if !exists {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}
		userClaim := user.(*dtos.UserClaim)
		if userClaim.Role != role {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not Authorized"})
			return
		}
		context.Next()
	}
}

// generate user token
func GenerateToken(user dtos.UserClaim) (string, error) {
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// Create the JWT claims, which includes the user ID, Role, Email and expiration time 7days
	token.Claims = jwt.MapClaims{
		"iss":   configs.JWTIssuer,
		"aud":   configs.JWTAudience,
		"ID":    user.ID,
		"Role":  user.Role,
		"Email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	return token.SignedString([]byte(configs.JWTSecret))
}
