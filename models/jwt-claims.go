package models

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	jwt.StandardClaims
	Id    int `json:"id"`
	FullName string `json:"fullname"`
	Status string `json:"status"`
}