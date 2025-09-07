package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/stormtomten/boot_dex/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	Cache      *pokecache.Cache
}

func NewClient(timeout, cacheTimeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
		baseURL:    baseURL,
		Cache:      pokecache.NewCache(cacheTimeout),
	}
}

func (c *Client) doGet(path string) ([]byte, error) {
	if !strings.HasPrefix(path, "https:") {
		path = c.baseURL + path
	}
	if b, ok := c.Cache.Get(path); ok {
		return b, nil
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.Cache.Add(path, b)
	return b, nil
}
