package repositories

import (
	"fmt"
	"os"
	"simple-catalog-v2/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func GenerateUserClaims(id int, fullname, status string) (models.UserClaims, error) {
	if err := godotenv.Load(); err != nil {
		return models.UserClaims{}, err
	}
	appName := os.Getenv("APPLICATION_NAME")
	claims := models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    appName,
			ExpiresAt: time.Now().Add(time.Duration(10) * time.Hour).Unix(),
		},
		Id:    id,
		FullName: fullname,
		Status: status,
	}
	return claims, nil
}

func GenerateUserToken(claims models.UserClaims) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	signature := []byte(os.Getenv("JWT_SIGNATURE_KEY"))
	signedToken, err := token.SignedString(signature)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GetUserClaims(tokenString string) (*models.UserClaims, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	signature := []byte(os.Getenv("JWT_SIGNATURE_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256{
			return nil, fmt.Errorf("signing method invalid")
		} 

		return signature, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func GetUserClaimsFromContext(ctx echo.Context) (*models.UserClaims) {
	userInfo := ctx.Get("userInfo").(*models.UserClaims)
	return userInfo
}

func GetJwtTokenFromCookies(ctx echo.Context) (string, bool) {
	value, err := GetCookie(ctx, "jwt")
	if err != nil {
		return "", false
	}

	return value, true
}