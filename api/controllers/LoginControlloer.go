package contorllers

import (
	"fmt"
	"net/http"
	"profileyou/api/LoCred"
	"profileyou/api/service"

	"github.com/gin-gonic/gin"
)

// login contorller interface
type LoginController interface {
	Authenticate(ctx *gin.Context)
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func (controller *loginController) Authenticate(ctx *gin.Context) {

	fmt.Println("LOGIN: ")
	token := controller.Login(ctx)
	fmt.Printf("Token: %v\n", token)
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}

}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential LoCred.LoginCredentials
	fmt.Println("LoginController login cunf run")
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}
