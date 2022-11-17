package fattyauth

import (
	"os"
	"time"

	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/token"
	log "github.com/go-pkgz/lgr"
)

func InitAuth() *auth.Service {
	options := auth.Opts{
		SecretReader: token.SecretFunc(func(token string) (string, error) {
			return "secret", nil
		}),
		TokenDuration: time.Minute * 5,
		CookieDuration: time.Hour * 24,
		DisableXSRF: true, // TODO: enable this
		Issuer: "fatty",
		URL: "http://localhost:3000",
		ClaimsUpd: token.ClaimsUpdFunc(func (claims token.Claims) token.Claims {
			if claims.User != nil && claims.User.Email == os.Getenv("ADMIN_EMAIL") {
				claims.User.SetAdmin(true)
			}

			return claims
		}),
		Logger: log.Default(),
		UseGravatar: false,
		AvatarStore: avatar.NewLocalFS("/tmp"),
	}

	service := auth.NewService(options)
	service.AddProvider("google", os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"))

	return service
}