package googleoauth2

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"pulsar/pkg/oauth2identity"
	"strings"
)

//1. Get code through consent screen
//2. Use code to get access token and id token
//3. Use access token and id token to get user info

const (
	TokenEndpoint    = "https://oauth2.googleapis.com/token"
	UserInfoEndpoint = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token"
	RedirectURI      = "http://localhost:3000/users/oauth/google"
)

// Token represent access token and id token
type Token struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	ExpiresIn   uint64 `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type UserInfo struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	ProfilePicture string `json:"picture"`
}

func (ui *UserInfo) GetName() (string, error) {
	return ui.Name, nil
}

func (ui *UserInfo) GetEmail() (string, error) {
	return ui.Email, nil
}

func (ui *UserInfo) GetProfilePicture() (string, error) {
	return ui.ProfilePicture, nil
}

// GetToken return an instance of Token
func GetToken(ctx context.Context, code string) (*Token, error) {
	client := http.DefaultClient
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("client_id", "1080567179937-2al8ksrahq6ags30gh8s969cgaivcj9f.apps.googleusercontent.com")
	form.Add("client_secret", "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")
	form.Add("redirect_uri", RedirectURI)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, TokenEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad request")
	}
	decoder := json.NewDecoder(res.Body)
	var token Token
	err = decoder.Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// GetUserInfo return an instance of UserInfo
func GetUserInfo(ctx context.Context, token *Token) (oauth2identity.Oauth2Identity, error) {
	client := http.DefaultClient
	endpoint := fmt.Sprintf("%s=%s", UserInfoEndpoint, token.AccessToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.IDToken))
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad request")
	}
	decoder := json.NewDecoder(res.Body)
	var userInfo UserInfo
	err = decoder.Decode(&userInfo)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
