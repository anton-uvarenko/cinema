package middleware

import (
	"github.com/anton-uvarenko/cinema/broker-service/internal/core"
	"github.com/anton-uvarenko/cinema/broker-service/internal/pkg"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	UserType []core.UserType
	Recovery bool
}

func (m AuthMiddleware) TokenVerify(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		token := strings.Split(header, " ")
		if len(token) < 2 {
			http.Error(w, "no jwt", http.StatusUnauthorized)
			return
		}

		if token[0] != "Bearer" {
			http.Error(w, "not bearer authorization", http.StatusUnauthorized)
			return
		}

		err := pkg.Verify(token[1], m.UserType, m.Recovery)

		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
