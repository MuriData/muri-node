package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/MuriData/muri-node/config"
)

// Client is an HTTP client for the Kubo IPFS API.
type Client struct {
	apiURL string
	http   *http.Client
}

// NewClient creates an IPFS client from config.
func NewClient(cfg config.IPFSConfig) *Client {
	timeout := cfg.Timeout.Duration
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	return &Client{
		apiURL: cfg.APIURL,
		http:   &http.Client{Timeout: timeout},
	}
}

// Cat fetches the raw bytes of a CID from IPFS.
func (c *Client) Cat(ctx context.Context, cid string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/v0/cat?arg=%s", c.apiURL, cid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ipfs cat: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ipfs cat %s: status %d: %s", cid, resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}

// addResponse is the JSON response from /api/v0/add.
type addResponse struct {
	Hash string `json:"Hash"`
	Name string `json:"Name"`
	Size string `json:"Size"`
}

// Add uploads data to IPFS and returns the CID.
func (c *Client) Add(ctx context.Context, data []byte, filename string) (string, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("create form file: %w", err)
	}
	if _, err := part.Write(data); err != nil {
		return "", fmt.Errorf("write data: %w", err)
	}
	writer.Close()

	url := fmt.Sprintf("%s/api/v0/add", c.apiURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &body)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("ipfs add: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ipfs add: status %d: %s", resp.StatusCode, string(respBody))
	}

	var result addResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	return result.Hash, nil
}

// Pin pins a CID to the local IPFS node.
func (c *Client) Pin(ctx context.Context, cid string) error {
	url := fmt.Sprintf("%s/api/v0/pin/add?arg=%s", c.apiURL, cid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("ipfs pin: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("ipfs pin %s: status %d: %s", cid, resp.StatusCode, string(body))
	}

	return nil
}

// IsPinned checks if a CID is pinned locally.
func (c *Client) IsPinned(ctx context.Context, cid string) (bool, error) {
	url := fmt.Sprintf("%s/api/v0/pin/ls?arg=%s", c.apiURL, cid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return false, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return false, fmt.Errorf("ipfs pin ls: %w", err)
	}
	defer resp.Body.Close()

	// Drain body for HTTP connection reuse
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	// Kubo returns 500 for "not pinned" — distinguish from other errors
	if resp.StatusCode == http.StatusInternalServerError {
		return false, nil
	}
	return false, fmt.Errorf("ipfs pin ls %s: status %d: %s", cid, resp.StatusCode, string(body))
}

// Ping checks connectivity to the IPFS node.
func (c *Client) Ping(ctx context.Context) error {
	url := fmt.Sprintf("%s/api/v0/id", c.apiURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("ipfs ping: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ipfs ping: status %d", resp.StatusCode)
	}

	return nil
}
