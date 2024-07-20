package handler

import (
	"BlogAPI/model"
	"BlogAPI/pkg/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReceiveUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(ctx *gin.Context) {
	var rUser ReceiveUser
	if err := ctx.ShouldBind(&rUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Couldn't read user's info",
		})
		return
	}
	var user model.User
	HashPassword, err := auth.HashPassword(rUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	err = user.Create(rUser.Username, HashPassword, rUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"Status": "Create user successfully",
	})
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var uLogin UserLogin
	if err := ctx.ShouldBind(&uLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("wrong json Login format"))
		return
	}
	var user model.User
	res, err := user.CheckIfExist("username", uLogin.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse(err.Error()))
		return
	}
	if !res {
		//ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("username isn't exist"))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   "This username already exists",
			"username": uLogin.Username,
		})
		return
	}
	user.UserName = uLogin.Username
	if err = user.GetUserByUsername(uLogin.Username); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("Sever Error"))
		return
	}
	if err = auth.ValidatePassword(user.Password, uLogin.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("Wrong Password!"))
		return
	}
	expireTime := time.Now().Add(time.Minute * 1).Unix()
	claims := auth.Claims{
		ID: *user.ID,
	}
	claims.StandardClaims.ExpiresAt = expireTime
	tokenString, err := auth.GenerateTokenString(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("Couldn't generated JWT Token"))
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"token": tokenString,
		"user":  user.UserName,
	})
}

func GetUserByUsername(ctx *gin.Context) {
	var uLogin UserLogin
	if err := ctx.ShouldBind(&uLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse("wrong json Login format"))
		return
	}
	var user model.User
	if err := user.GetUserByUsername(uLogin.Username); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"username": user.UserName,
		"ID":       user.ID,
	})
} /*
func GetUser(ctx *gin.Context) {
	token := auth.GetTokenString(ctx)
	claims, err := auth.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.ErrorResponse("Invalid token"))
		return
	}
	user := model.User{}
	res, err = user.CheckIfExist("id", strconv.claims.ID)
}
func GetUser(ctx *gin.Context) {
	token := auth.GetTokenString(ctx)
	claims, err := auth.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized,
			models.ErrorResponse("Invalid token"))
		return
	}

	user := models.User{}
	err = user.GetOne("_id", claims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			models.ErrorResponse("Fail to get user"))
		return
	}

	// Hide password
	user.Password = "*"

	ctx.JSON(http.StatusOK,
		models.SuccessResponse("Get user successfully", gin.H{"user": user}))
}
*/
