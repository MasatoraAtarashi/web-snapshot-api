package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL    = "http://localhost:3000/"
	apiVersion = "v1"
)

type Client struct {
	ContentService ContentService

	HTTPClient  *http.Client
	AccessToken string
	Client      string
	UID         string
	BaseURL     string
}

func NewClient(accessToken string, client string, uid string) *Client {
	var cli Client
	cli.ContentService = &contentService{cli: &cli}
	cli.AccessToken = accessToken
	cli.Client = client
	cli.UID = uid
	cli.BaseURL = baseURL + apiVersion
	return &cli
}

func (cli *Client) httpClient() *http.Client {
	if cli.HTTPClient != nil {
		return cli.HTTPClient
	}
	return http.DefaultClient
}

func (cli *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	req.Header.Set("access-token", cli.AccessToken)
	req.Header.Set("client", cli.Client)
	req.Header.Set("uid", cli.UID)
	return cli.httpClient().Do(req)
}

func (cli *Client) putForm(ctx context.Context, path string, data url.Values, v interface{}) error {
	reqURL := cli.BaseURL + "/" + path
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest(http.MethodPut, reqURL, body)
	if err != nil {
		return fmt.Errorf("HTTP request Error: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := cli.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode > -http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return cli.error(resp.StatusCode, resp.Body)
	}

	if v == nil {
		return nil
	}

	var respBody io.Reader = resp.Body
	if err := json.NewDecoder(respBody).Decode(v); err != nil {
		return fmt.Errorf("Response parse Error: %w", err)
	}

	return nil
}

func (cli *Client) error(statusCode int, body io.Reader) error {
	var aerr APIError
	if err := json.NewDecoder(body).Decode(&aerr); err != nil {
		return &APIError{StatusCode: statusCode}
	}
	aerr.StatusCode = statusCode
	return &aerr
}
