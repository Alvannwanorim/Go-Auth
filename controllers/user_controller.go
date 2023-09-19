package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alvannwanorim/go-auth/initializers"
	"github.com/alvannwanorim/go-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JwtClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateUser(c *gin.Context) {
	var existingUser models.User
	var body struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})

		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request failed to process",
		})
		return
	}

	//find existing user
	initializers.DB.First(&existingUser, "email=?", body.Email)
	if existingUser.ID != 0 || existingUser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user already exists",
		})
		return
	}

	//Create new user
	user := models.User{Email: body.Email, FirstName: body.FirstName, LastName: body.LastName, Password: string(password)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(c *gin.Context) {

	JWT_SECRET := os.Getenv("JWT_SECRET")

	var user models.User

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error validating request body",
		})
		return
	}

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 || user.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or password mismatch",
		})
		return
	}
	claim := JwtClaim{
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		fmt.Println(tokenString, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to process request",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}

// Get Logged in User
func Validate(c *gin.Context) {
	email, _ := c.Get("user")

	user := GetUserByEmail(email)

	if user.ID == 0 || user.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": user,
	})

}

func GetUserByEmail(email any) *models.User {
	user := new(models.User)
	initializers.DB.First(&user, "email = ?", email)
	return user
}

func GetUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Find(&users)
	fmt.Println(result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
