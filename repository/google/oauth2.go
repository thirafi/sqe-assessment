package google

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "1063922470623-9lva3uhgqobftft2kt7ornetidi43n9u.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

type (
	GoogleUserResult struct {
		Id             string
		Name           string
		Email          string
		Verified_email bool
		Picture        string
	}

	GoogleOauthToken struct {
		Email    string
		ExpireIn float64
	}
)

func GetGoogleUser(AccessToken string) (*GoogleUserResult, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", AccessToken)

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", IdToken))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return nil, err
	}

	userBody := &GoogleUserResult{
		Id:             GoogleUserRes["id"].(string),
		Email:          GoogleUserRes["email"].(string),
		Name:           GoogleUserRes["name"].(string),
		Verified_email: GoogleUserRes["verified_email"].(bool),
		Picture:        GoogleUserRes["picture"].(string),
	}

	return userBody, nil
}

func GetGoogleOauthToken(ctx *gin.Context, code string) (tokenBody *oauth2.Token, err error) {
	tokenBody, err = googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "could not exchange oauth code")
		return
	}

	if !tokenBody.Valid() {
		ctx.JSON(http.StatusInternalServerError, "invalid access token")
		return nil, fmt.Errorf("invalid access token")
	}

	return
}
