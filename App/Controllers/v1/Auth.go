package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func IndexLogin(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			Audience:  "1",
			ExpiresAt: time.Now().Unix(),
			Issuer:    "index.auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	c.JSON(http.StatusOK, gin.H{
		"name":  "IndexLogin",
		"token": token,
		"ss":    ss,
	})
}
