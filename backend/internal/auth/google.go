package auth

import (
	"context"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuth struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
	Endpoint     oauth2.Endpoint
}

func CreateGoogleAuth(clientID, clientSecret, redirectURL string) GoogleAuth {
	return GoogleAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	// return oauth2.Config{
	// RedirectURL:  "http://localhost:3000/auth/google/callback",
	// ClientID:     "54384433304-cq2q6nh3hukkf5b2s0ppsmpos6rn9i9h.apps.googleusercontent.com",
	// ClientSecret: "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",
	// Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
	// 	"https://www.googleapis.com/auth/userinfo.profile"},
	// Endpoint: google.Endpoint,
	// }
}

func (g *GoogleAuth) GetConfig() oauth2.Config {
	return oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		RedirectURL:  g.RedirectURL,
		Scopes:       g.Scopes,
		Endpoint:     g.Endpoint,
	}
}

func (g *GoogleAuth) VerifyCallBack(ctx context.Context, code string) (string, error) {
	googleCon := g.GetConfig()

	token, err := googleCon.Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	userData, err := g.fetchUserInfo(token)
	if err != nil {
		return "", err
	}

	return userData, nil
}

func (g *GoogleAuth) fetchUserInfo(token *oauth2.Token) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return "", err
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(userData), nil
}
