package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// HTTPResponse はHTTPリクエストの結果を表す構造体です
type HTTPResponse struct {
	StatusCode int
	Body       []byte
	Header     http.Header
}

// HTTPRequestsUtils インターフェースはHTTPリクエストのメソッドを定義します
type HTTPRequestsUtils interface {
	// Get は指定されたURLにGETリクエストを送信し、HTTPResponseを返します
	Get(url string) (*HTTPResponse, error)
	// Post は指定されたURLにPOSTリクエストを送信し、HTTPResponseを返します
	Post(url string, contentType string, body []byte) (*HTTPResponse, error)
	// ここに他のHTTPメソッド（PUT, DELETE など）を追加できます
}

// httpRequestsUtils は HTTPRequestsUtils インターフェースの実装構造体です
type httpRequestsUtils struct {
	client *http.Client
}

// NewHTTPRequestsUtils は新しい HTTPRequestsUtils インスタンスを作成します
func NewHTTPRequestsUtils() HTTPRequestsUtils {
	return &httpRequestsUtils{
		client: &http.Client{},
	}
}

// Get メソッドの実装
func (h *httpRequestsUtils) Get(url string) (*HTTPResponse, error) {
	resp, err := h.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return &HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       body,
		Header:     resp.Header,
	}, nil
}

// Post メソッドの実装
func (h *httpRequestsUtils) Post(url string, contentType string, body []byte) (*HTTPResponse, error) {
	resp, err := h.client.Post(url, contentType, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return &HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       respBody,
		Header:     resp.Header,
	}, nil
}

// 他のHTTPメソッドの実装をここに追加できます
// 例: Put, Delete など
