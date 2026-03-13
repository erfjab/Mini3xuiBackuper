package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type Client struct {
	host string
	http *http.Client
}

type loginResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func New(host string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("create cookie jar: %w", err)
	}
	return &Client{
		host: strings.TrimRight(host, "/"),
		http: &http.Client{Jar: jar},
	}, nil
}

func (c *Client) Login(username, password string) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", username)
	_ = writer.WriteField("password", password)
	_ = writer.WriteField("twoFactorCode", "")
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, c.host+"/login", payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var result loginResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return fmt.Errorf("decode login response: %w", err)
	}
	if !result.Success {
		return fmt.Errorf("login failed: %s", result.Msg)
	}
	return nil
}

func (c *Client) DownloadDB() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.host+"/server/getDb", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	return io.ReadAll(res.Body)
}
