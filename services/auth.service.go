package services

import (
	"fmt"
	"lfs-portal/enums"
	"lfs-portal/models"
	"lfs-portal/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	Email    string         `json:"email"`
	UserType enums.UserType `json:"userType"`
	jwt.StandardClaims
}

var jwtKey = []byte(utils.Config.JWT_SECRET)

func GenerateJWT(user models.AuthUser) (string, error) {
	// expire after two hours
	expirationTime := time.Now().Add(120 * time.Minute)
	claims := &Claims{
		Email:    user.Email,
		UserType: user.UserType,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		logrus.Error(fmt.Sprintf("[GenerateJWT] Error generating JSON Web Token, %v", err))
		return "", err
	}
	return tokenString, nil
}

func IsAuthenticated(token string) bool {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return false
	}

	return true
}

func IsAuthorized(token string, resource interface{}, action enums.Action) bool {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return false
	}

	if claims.UserType == enums.Admin {
		return true
	}

	if action == enums.Delete {
		// admin and customer admin

	}

	if action == enums.Execute {

	}
	return true
}

func AuthenticationMiddleware(c *gin.Context) {
	if c.FullPath() == "/api/user/authenticate" {
		c.Next()
		return
	}

	token, err := c.Cookie("lfs-auth-token")
	if err != nil {
		logrus.Error(fmt.Sprintf("[AuthenticationMiddleware] Error getting cookie, %v", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetUnauthorizedError())
		return
	}

	if token == "" {
		logrus.Error("[AuthenticationMiddleware] Cookie token missing.")
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetUnauthorizedError())
		return
	}

	if !IsAuthenticated(token) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetUnauthorizedError())
		return
	}
	c.Next()
}
