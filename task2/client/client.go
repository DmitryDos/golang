package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"task2/dto"
	"time"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c *Client) GetVersion() ([]byte, error) {
	req, err := http.NewRequest("GET", c.url+"/version", nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) GetHardOp() (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", c.url+"/hard-op", nil)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return false, http.StatusInternalServerError, nil
		}
		return false, http.StatusInternalServerError, err
	}
	defer response.Body.Close()
	return true, response.StatusCode, nil
}

func (c *Client) PostDecode(inputString string) (string, error) {
	reqBody, err := json.Marshal(dto.Request{Body: inputString})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.url+"/decode", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var decoded dto.Response
	err = json.Unmarshal(body, &decoded)
	if err != nil {
		return "", err
	}
	return decoded.Body, nil
}
