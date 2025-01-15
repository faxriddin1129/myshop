package middleware

import (
	"myshop/models"
	"myshop/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, map[string]string{
				"msg": "Authorization header missing",
			})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		userID, valid := models.TokenExists(tokenString)
		if !valid {
			utils.RespondWithError(w, http.StatusUnauthorized, map[string]string{
				"msg": "Invalid or expired token",
			})
			return
		}

		ctx := utils.SetUserIDToContext(r.Context(), userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
