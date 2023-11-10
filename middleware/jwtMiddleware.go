package middleware

import (
	"SQ-Assessment/constant"
	"SQ-Assessment/repository/database"
	"SQ-Assessment/repository/google"
	"SQ-Assessment/util"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	userRepository database.UserRepository
}

func NewMiddleware(
	userRepository database.UserRepository,
) *Middleware {
	return &Middleware{
		userRepository,
	}
}

func (m *Middleware) CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(constant.SECRET_JWT)
	return token.SignedString(byteSecret)
}

func (m *Middleware) ExtractTokenUserId(e *gin.Context) int {
	user, _ := e.Get("user")
	if user.(*jwt.Token).Valid {
		claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}

	return 0
}

func (m *Middleware) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("AccessToken")
		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, util.Error{
				Code:    http.StatusForbidden,
				Message: "No Authorization header provided",
			})
			return
		}

		user, err := google.GetGoogleUser(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.Error{
				Code:    http.StatusUnauthorized,
				Message: "Failed to get user from google",
			})
			return
		}

		validUser, errUser := m.userRepository.GetUserByEmail(user.Email)
		if errUser.NotNil() {
			return
		}
		ctx.Set("user", validUser)
		ctx.Next()
	}
}
