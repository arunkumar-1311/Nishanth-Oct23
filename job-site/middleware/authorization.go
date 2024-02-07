package middleware

import (
	"context"

	"net/http"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"
)

type Authorization interface {
	Authorization(kithttp.DecodeRequestFunc) kithttp.DecodeRequestFunc
}


func (svc Middleware) Authorization(decode kithttp.DecodeRequestFunc) kithttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		token := r.Header.Get("Authorization")
		if token == "" {
			data := map[string]interface{}{
				"Regiter": "localhost:8000/register",
				"Login":   "localhost:8000/login",
			}
			return data, nil
		}

		if _, err = svc.VerifyToken(token[7:]); err != nil {
			return err, nil
		}

		if err := svc.AdminAccess(token); strings.Contains(r.URL.Path, "admin") && err != nil {
			return err, nil
		}
		return decode(ctx, r)
	}
}
