package jwt

import (
	"strings"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
)

func JWTAuth(h engine.Handle) engine.Handle {
	return func(c *engine.Context) (interface{}, error) {
		authHeader := c.Context.Request.Header.Get("Authorization")
		if authHeader == "" {
			return nil, e.ERR_AUTH_EMPTY
		}
		// Bearer
		headers := strings.SplitN(authHeader, " ", 2)
		if len(headers) != 2 || headers[0] != "Bearer" {
			return nil, e.ERR_AUTH_FORMAT
		}

		tokenString, err := ParseToken(headers[1])
		if err != nil {
			return nil, err
		}

		c.Context.Set("token", tokenString)

		if r, err := h(c); err == nil {
			return r, nil
		} else {
			return nil, err
		}
	}
}
