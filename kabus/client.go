package kabus

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// client - HTTPクライアント
type client struct {
	url   string // URL
	token string // リクエストトークン
}

// get - GETリクエスト
func (c *client) get(ctx context.Context, pathParam string, queryParam string) (int, []byte, error) {
	u, err := url.Parse(c.url)
	if err != nil {
		return 0, nil, err
	}
	if pathParam != "" {
		u.Path += "/" + pathParam
	}
	if queryParam != "" {
		u.RawQuery = queryParam
	}

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, nil, err
	}
	if c.token != "" {
		req.Header.Set("X-API-KEY", c.token)
	}

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, b, nil
}

// post - POSTリクエスト
func (c *client) post(ctx context.Context, request []byte) (int, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", c.url, bytes.NewReader(request))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if c.token != "" {
		req.Header.Set("X-API-KEY", c.token)
	}

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, b, nil
}

// put - PUTリクエスト
func (c *client) put(ctx context.Context, request []byte) (int, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, "PUT", c.url, bytes.NewReader(request))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if c.token != "" {
		req.Header.Set("X-API-KEY", c.token)
	}

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, b, nil
}

// parseResponse - レスポンスをパースする
func parseResponse(code int, body []byte, v interface{}) error {
	if code == http.StatusOK {
		if err := json.Unmarshal(body, v); err != nil {
			return err
		}
		return nil
	} else {
		var errRes ErrorResponse
		if err := json.Unmarshal(body, &errRes); err != nil {
			return err
		}
		errRes.StatusCode = code
		errRes.Body = string(body)
		return errRes
	}
}

// createURL - リクエスト先のURLを生成する
func createURL(path string, isProd bool) string {
	host := "localhost:18081"
	if isProd {
		host = "localhost:18080"
	}
	return "http://" + host + strings.ReplaceAll("/kabusapi/"+path, "//", "/")
}
