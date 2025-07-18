package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("SECRET_KEY_KAMU") // ganti dengan secret dari env/config
// ValidateJWT untuk mengecek dan mengambil user_id dari token
func ValidateJWT(tokenStr string) (uint, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        // Validasi algoritma
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtKey, nil
    })

    if err != nil {
        return 0, err
    }

    // Ambil claims dan cek valid
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if userIDFloat, ok := claims["user_id"].(float64); ok {
            return uint(userIDFloat), nil
        }
        return 0, errors.New("user_id tidak ditemukan dalam token")
    }

    return 0, errors.New("token tidak valid")
}
