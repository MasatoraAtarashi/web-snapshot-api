package ws

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AuthService interface {
	SignIn(ctx context.Context, params *AuthParams) error
}

type authService struct {
	cli *Client
}

func (s *authService) SignIn(ctx context.Context, params *AuthParams) error {
	// NOTE: うまく抽象化したいけど、妥協
	path := "auth/sign_in"
	data := url.Values{}
	data.Set("email", params.EMail)
	data.Set("password", params.PassWord)

	reqURL := s.cli.BaseURL + "/" + path
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest(http.MethodPost, reqURL, body)
	if err != nil {
		return fmt.Errorf("HTTP request Error: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)
	resp, err := s.cli.httpClient().Do(req)
	if err != nil {
		return err
	}
	if !(resp.StatusCode > -http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return s.cli.error(resp.StatusCode, resp.Body)
	}

	token := resp.Header.Get("access-token")
	client := resp.Header.Get("client")
	uid := resp.Header.Get("uid")
	if token == "" || client == "" || uid == "" {
		return fmt.Errorf("Error: cannnot get auth header from response")
	}

	s.cli.AuthHeader = &AuthHeader{
		AccessToken: token,
		Client:      client,
		UID:         uid,
	}
	return nil
}
