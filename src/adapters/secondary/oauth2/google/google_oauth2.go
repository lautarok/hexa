package google

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/lautarok/hexa/src/core/ports"
	"github.com/lautarok/hexa/src/core/ports/dtos"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOauth2Adapter struct {
	config *oauth2.Config
}

type GoogleOauth2AdapterDeps struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func NewGoogleOauth2Adapter(deps *GoogleOauth2AdapterDeps) ports.GoogleOAuth2Port {
	return &GoogleOauth2Adapter{
		config: &oauth2.Config{
			ClientID:     deps.ClientID,
			ClientSecret: deps.ClientSecret,
			RedirectURL:  deps.RedirectURL,
			Scopes:       []string{"openid", "email", "profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (adapter *GoogleOauth2Adapter) GetLoginURL(ctx context.Context) string {
	return adapter.config.AuthCodeURL("state")
}

func (adapter *GoogleOauth2Adapter) HandleCallback(ctx context.Context, code string) (*dtos.GoogleIdentityDto, error) {
	oauth2Token, err := adapter.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := adapter.config.Client(ctx, oauth2Token)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var userInfo struct {
		ID         string `json:"id"`
		Email      string `json:"email"`
		Name       string `json:"name"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
	}

	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &dtos.GoogleIdentityDto{
		GoogleID:   userInfo.ID,
		Email:      userInfo.Email,
		Name:       userInfo.Name,
		GivenName:  userInfo.GivenName,
		FamilyName: userInfo.FamilyName,
	}, nil
}

func (adapter *GoogleOauth2Adapter) IsInvalidGrantError(err error) bool {
	var rErr *oauth2.RetrieveError
	if errors.As(err, &rErr) {
		return strings.Contains(string(rErr.Body), "invalid_grant")
	}
	return false
}
