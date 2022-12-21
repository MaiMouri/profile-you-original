package contorllers

import (
	"fmt"
	"net/http"
	"os"
	"profileyou/api/LoCred"
	"profileyou/api/domain/model/user"
	"profileyou/api/service"
	"profileyou/api/usecase"
	"profileyou/api/utils/errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// login contorller interface
type UserController interface {
	Authenticate(ctx *gin.Context)
	Login(ctx *gin.Context) string
	Signup(ctx *gin.Context)
}

type userController struct {
	loginService  service.LoginService
	signupService service.SignupService
	jWtService    service.JWTService
	userUseCase   usecase.UserUseCase
}

// likes to Usecase by "ku"
func NewUserController(uu usecase.UserUseCase) userController {
	return userController{
		userUseCase: uu,
	}

}

func (controller *userController) Authenticate(ctx *gin.Context) {

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

func UserHandler(loginService service.LoginService,
	jWtService service.JWTService) UserController {
	return &userController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *userController) Login(ctx *gin.Context) string {

	var credential LoCred.LoginCredentials
	fmt.Println("LoginController login cunf run")
	err := ctx.ShouldBindJSON(&credential)
	if err != nil {
		return ""
	}

	user, err := controller.userUseCase.GetUserForAuth(credential.Email)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad Request")
		ctx.IndentedJSON(apiErr.Status, apiErr)
		return ""
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//encoded string
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		panic(err)
	}

	return t

}

func (controller *userController) Signup(ctx *gin.Context) {
	var credential LoCred.LoginCredentials

	if err := ctx.ShouldBindJSON(&credential); err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad Request")
		ctx.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(credential.Password), 12)

	u := user.User{Email: credential.Email, Password: string(hash)}

	fmt.Printf("user Information %v, email %v, pass %v\n", u, u.Email, u.Password)
	u.Create()
	// err := controller.userUsecase.CreateUser(u.Email, u.Password)
	// if err != nil {
	// 	fmt.Println("ERROR OCCURED HERE")
	// 	apiErr := errors.InternalSeverError("Server Error when posting")
	// 	ctx.IndentedJSON(apiErr.Status, apiErr)
	// 	fmt.Println(err)
	// }

	fmt.Println(u.Email, u.Password)

	ctx.JSON(200, gin.H{
		// "token": str_out,
		"message": "Register success!",
	})
}
