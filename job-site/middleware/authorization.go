package middleware

import (
	"context"

	"net/http"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"
)

// Helps to check verify the token and authorize the admin
func (svc Middleware) Authorization(decode kithttp.DecodeRequestFunc) kithttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (request interface{}, err error) {

		UUID := r.Header.Get("Authorization")
		if UUID == "" {
			data := map[string]interface{}{
				"Regiter": "localhost:8000/register",
				"Login":   "localhost:8000/login",
			}
			return data, nil
		}

		token, err := svc.DB.GetRedisCache(UUID[7:])
		if err != nil {
			return "Invalid Token Please Login", nil
		}

		if _, err = svc.VerifyToken(token[7:]); err != nil {
			return err.Error(), nil
		}

		if err := svc.AdminAccess(token); strings.Contains(r.URL.Path, "/admin/") && err != nil {
			return err.Error(), nil
		}
		r.Header.Set("Authorization", token)
		return decode(ctx, r)
	}
}
