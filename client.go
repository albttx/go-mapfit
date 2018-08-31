package mapfit

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type apiConfig struct {
	method  string
	baseUrl string
	path    string
}

type Client struct {
	httpClient *http.Client
	token      string
}

func NewClient(token string) *Client {
	client := &Client{
		token: token,
	}
	return client
}

func (c *Client) request(cfg *apiConfig, bodyReq io.Reader) ([]byte, error) {
	req, err := http.NewRequest(
		cfg.method,
		fmt.Sprintf("%s/%s?api_key=%s", cfg.baseUrl, cfg.path, c.token),
		bodyReq,
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyResp, err
}
