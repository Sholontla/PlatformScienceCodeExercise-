package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyQWRtaW4yIiwiZXhwIjoxNjc5MjIwNjQzLCJTY29wZSI6ImFkbWluIiwicm9sZXMiOlsiYWRtaW4iXSwicGVybWlzc2lvbnMiOlsiY3JlYXRlX2ludmVudG9yeSIsImRlbGV0ZV9pbnZlbnRvcnkiLCJyZWFkX2ludmVudG9yeSIsInVwZGF0ZV9pbnZlbnRvcnkiXX0.FxelUzIBQwfQ3CYiQBN5kDnmCcKTArbWx9jPUvsfedA"
	key := []byte("9ikNKMPJQ8dNIsDKQ9zaGbjH4Zp5IU0NA==")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key used to sign the token
		return key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("Invalid signature")
		} else {
			fmt.Println("Error parsing token:", err)
		}
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token claims:", claims)
	} else {
		fmt.Println("Invalid token")
	}
}
