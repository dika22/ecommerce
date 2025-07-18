package middleware

import (
	"context"
	"net/http"
	"order-service/package/utils"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
        userID, err := utils.ValidateJWT(tokenStr) // validasi lokal
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Simpan user ID ke context
        ctx := context.WithValue(r.Context(), "user_id", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
