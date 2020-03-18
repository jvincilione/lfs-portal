package services

import (
	"fmt"
	"lfs-portal/enums"
	"lfs-portal/models"
	"lfs-portal/utils"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	ID        uint           `json:"id"`
	Email     string         `json:"email"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	UserType  enums.UserType `json:"userType"`
	Companies []uint         `json:"companies"`
	CompanyID uint           `json:"companyId"`
	jwt.StandardClaims
}

var jwtKey = []byte(utils.Config.JWT_SECRET)

func GenerateJWT(user models.AuthUser) (string, error) {
	// expire after two hours
	expirationTime := time.Now().Add(120 * time.Minute)
	claims := &Claims{
		ID:        user.ID,
		Email:     user.Email,
		UserType:  user.UserType,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Companies: user.Companies,
		CompanyID: user.CompanyID,
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
		// admin and company admin

	}

	if action == enums.Execute {

	}
	return true
}

func LogoutUser(c *gin.Context) {
	c.SetCookie("lfs-auth-token", "", 0, "/", utils.Config.DOMAIN, http.SameSiteNoneMode, false, false)
}

func AuthenticationMiddleware(c *gin.Context) {
	logrus.Error(c.FullPath())
	if c.FullPath() == "/api/user/authenticate" || c.FullPath() == "/api/user/logout" || !strings.HasPrefix(c.FullPath(), "/api/") {
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
		LogoutUser(c)
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetUnauthorizedError())
		return
	}

	if c.Request.Method == "POST" || c.Request.Method == "PATH" || c.Request.Method == "DELETE" {
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "You are not authorized - no bearer",
			})
			return
		}
		splitToken := strings.Split(bearerToken, "Bearer")
		bearerToken = strings.Trim(splitToken[1], " ")
		logrus.Info(bearerToken)
		logrus.Info(token)
		if bearerToken != token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "You are not authorized - bearer doesn't match",
			})
			return
		}
	}
	c.Next()
}
