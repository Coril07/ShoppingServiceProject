package user

import (
	"app/server/config"
	"app/server/domain/user"
	"app/server/utils/api_helper"
	jwtHelper "app/server/utils/jwt"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService *user.Service
	appConfig   *config.Configuration
}

// 实例化
func NewUserController(service *user.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		userService: service,
		appConfig:   appConfig,
	}
}

func (c *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	fmt.Printf("req: %v\n", req)
	newUser := user.NewUser(req.Username, req.Password, req.Password2, req.Gender, req.Birth, req.Url, req.Email, req.Address)
	err := c.userService.Create(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateUserResponse{
			Username: req.Username,
		})
}

func (c *Controller) CookieLogin(g *gin.Context) {
	isadmin, _ := g.Get("isadmin")
	val, t := isadmin.(bool)
	if t {
		g.JSON(
			http.StatusOK, LoginResponse{IsAdmin: val})
	} else {
		api_helper.HandleError(g, api_helper.ErrServerBug)
	}

}

func (c *Controller) Log_out(g *gin.Context) {
	g.SetCookie("Token", "", -1, "/", "localhost", false, true)
	g.SetCookie("Username", "", -1, "/", "localhost", false, true)
	session := sessions.Default(g)
	session.Clear()
	g.JSON(http.StatusOK, "")
}

func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	currentUser, err := c.userService.GetUser(req.Username, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	decodedClaims := jwtHelper.VerifyToken(currentUser.Token, c.appConfig.SecretKey)
	if decodedClaims == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp": time.Now().Add(
					24 *
						time.Hour).Unix(),
				"isAdmin": currentUser.IsAdmin,
			})
		token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.SecretKey)
		currentUser.Token = token
		err = c.userService.UpdateUser(&currentUser)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
	}
	g.SetCookie("Token", currentUser.Token, 60*60, "/", "localhost", false, true)
	g.SetCookie("Username", currentUser.Username, 60*60, "/", "localhost", false, true)
	session := sessions.Default(g)
	session.Set(currentUser.Token, currentUser.Username)
	g.JSON(
		http.StatusOK, LoginResponse{Username: currentUser.Username, UserId: currentUser.ID, IsAdmin: currentUser.IsAdmin})
}

// 验证token
func (c *Controller) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.SecretKey)

	g.JSON(http.StatusOK, decodedClaims)

}
