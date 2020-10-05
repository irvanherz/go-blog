package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/irvanherz/goblog/model"
	"github.com/irvanherz/goblog/service"
	"golang.org/x/crypto/bcrypt"
)

//AuthController hehehe
type AuthController struct {
	us *service.UserService
}

// NewAuthController heheh
func NewAuthController(us *service.UserService) *AuthController {
	return &AuthController{us: us}
}

// Login Auths
func (c *AuthController) Login(ctx *gin.Context) {
	var loginPayload model.LoginPayload

	if err := ctx.ShouldBind(&loginPayload); err != nil || loginPayload.Email == "" || loginPayload.Password == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Status: http.StatusBadRequest,
			Error:  model.NewRequestError("1234", err.Error()),
		})
		return
	}

	user, err := c.us.ReadByEmail(loginPayload.Email)
	if err != nil || user == nil {
		ctx.AbortWithStatusJSON(http.StatusOK, model.Response{
			Status: http.StatusOK,
			Error:  model.NewRequestError("auth fail", "Invalid password"),
		})
		return
	}
	if hashError := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(loginPayload.Password)); hashError != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, model.Response{
			Status: http.StatusOK,
			Error:  model.NewRequestError("auth fail", "Invalid password"),
		})
		return
	}
	authData := model.AuthData{}
	authData.UserID = user.ID
	authData.ExpiresAt = time.Now().Add(time.Minute * 60).Unix()
	authData.Issuer = "Goblog"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authData)

	// Sign and get the complete encoded token as a string using the secret
	generatedToken, _ := token.SignedString([]byte("DEFILATIFAH"))

	user.Password = nil

	ctx.JSON(http.StatusOK, model.Response{
		Status: http.StatusOK,
		Data: model.LoginResponsePayload{
			Token:    generatedToken,
			UserData: *user,
		},
	})
}
