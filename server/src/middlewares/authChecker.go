package middlewares

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"server/src/lib"
	"strings"
)

func validateToken(ctx context.Context, tok string, oidcConfig *lib.OIDCConfig) (*oidc.IDToken, error) {
	provider, err := oidc.NewProvider(ctx, oidcConfig.Issuer)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create oidc provider")
	}

	var verifier = provider.Verifier(&oidc.Config{
		ClientID: oidcConfig.Audience,
	})
	s, err := verifier.Verify(ctx, tok)
	if err != nil {
		return nil, errors.Wrap(err, "Could not validate token")
	}
	return s, nil
}

func parseTokenFromAuthHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("Missing authorization header")
	}

	splittedHeader := strings.Split(authHeader, " ")
	if len(splittedHeader) != 2 {
		return "", errors.New("Malformed Authorization header")
	}
	if splittedHeader[0] != "Bearer" {
		return "", errors.New("Server only accepts Bearer tokens")
	}

	return splittedHeader[1], nil
}

func AuthChecker(oidcConfig *lib.OIDCConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok, err := parseTokenFromAuthHeader(c)
		if err != nil {
			lib.ErrorMessage(c, err, http.StatusUnauthorized, "Invalid auhotization header")
			return
		}

		token, err := validateToken(c, tok, oidcConfig)
		if err != nil {
			lib.ErrorMessage(c, err, http.StatusUnauthorized, "Invalid token")
			return
		}

		var claims map[string]interface{}
		err = token.Claims(&claims)
		if err != nil {
			lib.ErrorMessage(c, err, http.StatusUnauthorized, "Could not parse token claims")
			return
		}

		c.Set("name", claims["name"])
		c.Set("userID", claims["sub"])
		c.Next()
	}
}
