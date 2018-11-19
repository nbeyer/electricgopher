package api

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

type AccessTokenInput struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

func (ati *AccessTokenInput) ToUrlValues() url.Values {
	result := url.Values{}
	result.Set("grant_type", ati.GrantType)
	result.Set("client_id", ati.ClientId)
	result.Set("client_secret", ati.ClientSecret)
	if ati.Email != "" {
		result.Set("email", ati.Email)
	}
	if ati.Password != "" {
		result.Set("password", ati.Password)
	}
	if ati.RefreshToken != "" {
		result.Set("refresh_token", ati.RefreshToken)
	}
	return result
}

type AccessTokenOutput struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    int    `json:"created_at"`
}

func (ato *AccessTokenOutput) String() string {
	return fmt.Sprintf(
		"{AccessToken:%s, TokenType:%s, ExpiresIn:%d, RefreshToken:%s, CreatedAt:%d}",
		ato.AccessToken,
		ato.TokenType,
		ato.ExpiresIn,
		ato.RefreshToken,
		ato.CreatedAt,
	)
}

func NewAccessTokenInputWithPassword(clientId, clientSecret, email, password string) *AccessTokenInput {
	return &AccessTokenInput{
		GrantType:    "password",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Email:        email,
		Password:     password,
	}
}

func NewAccessTokenInputWithRefreshToken(clientId, clientSecret, refreshToken string) *AccessTokenInput {
	return &AccessTokenInput{
		GrantType:    "refresh_token",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	}
}

func (c *Client) GetAccessToken(in *AccessTokenInput) (*AccessTokenOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.GetAccessToken(): begin")
	url := c.resolveUrl("/oauth/token")
	res, err := c.httpClient.PostForm(url, in.ToUrlValues())
	if err != nil {
		c.logger.Debugf("electricgopher.api.Client.GetAccessToken(): error posting request - %v", err)
		return nil, err
	}
	defer res.Body.Close()

	c.logger.Debugf("electricgopher.api.Client.GetAccessToken(): http response status - %s", res.Status)
	// check for 200
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to acquire token; response status %s", res.Status)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var ato AccessTokenOutput
	mustDeserializeJson(resBody, &ato)
	return &ato, nil
}
