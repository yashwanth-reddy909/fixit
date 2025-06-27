package http

import (
	"net/http"

	"fixit.com/backend/src/models/dto"
	"fixit.com/backend/src/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserSvc
}

func NewUserController(userService service.UserSvc) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Signup(ctx *gin.Context) {
	var req dto.SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.userService.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (c *UserController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userService.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("access_token", token, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (c *UserController) GoogleAuthCallback(ctx *gin.Context) {
	var req dto.GoogleAuthCallback
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userService.GoogleAuthCallback(ctx, req.Code, req.State)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("access_token", token, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Google Auth successful"})
}
