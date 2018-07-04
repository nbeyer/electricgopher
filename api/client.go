package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	Credentials   *Credentials
	logger        *logrus.Logger
	baseUri       string
	httpClient    *http.Client
	currentTokens *AccessTokenOutput
}

type Credentials struct {
	ClientId     string
	ClientSecret string
	Email        string
	Password     string
}

func NewClient(clientId, clientSecret, email, password, baseUri string, logger *logrus.Logger) *Client {
	if logger == nil {
		logger = logrus.StandardLogger()
	}
	c := &Client{
		Credentials: &Credentials{
			ClientId:     clientId,
			ClientSecret: clientSecret,
			Email:        email,
			Password:     password,
		},
		baseUri:    baseUri,
		httpClient: newHttpClient(logger),
		logger:     logger,
	}
	return c
}

func (c *Client) resolveUrl(path string) string {
	return fmt.Sprintf("%s%s", c.baseUri, path)
}

func (c *Client) acquireAccessToken() (string, error) {
	// TODO: add checking expiration and acquiring token via refresh
	if c.currentTokens != nil {
		return c.currentTokens.AccessToken, nil
	}
	err := c.Authenticate()
	if err != nil {
		return "", err
	}
	return c.currentTokens.AccessToken, nil
}

func (c *Client) doGet(path string, out interface{}) error {
	url := c.resolveUrl(path)
	c.logger.Debugf("tesla.electricgopher.api.Client.doGet(): GET %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	accessToken, err := c.acquireAccessToken()
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	res, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Debugf("tesla.electricgopher.api.Client.doGet(): error making request - %s", err.Error())
		return err
	}
	// check for 200
	if res.StatusCode != 200 {
		return fmt.Errorf("Unable to get resource; HTTP response status %s", res.Status)
	}
	// TODO: make this more resilient
	ct := res.Header["Content-Type"][0]
	if !strings.Contains(ct, "application/json") {
		return fmt.Errorf("Unable to parse response; unexpected content type - %s", ct)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Unable to read response body - %s", err.Error())
	}
	c.logger.Debugf("tesla.electricgopher.api.Client.doGet(): HTTP Response Body - %s", makeJsonPretty(resBody))
	// TODO: inline deserialization and add informative logging
	mustDeserializeJson(resBody, out)
	return nil
}

func (c *Client) Authenticate() error {
	ati := NewAccessTokenInputWithPassword(
		c.Credentials.ClientId,
		c.Credentials.ClientSecret,
		c.Credentials.Email,
		c.Credentials.Password,
	)
	ato, err := c.GetAccessToken(ati)
	if err != nil {
		return err
	}
	c.currentTokens = ato
	return nil
}
