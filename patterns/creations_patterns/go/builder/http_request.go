package builder

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HttpRequest represents an immutable HTTP request that can be executed
type HttpRequest struct {
	url         string
	method      string
	headers     map[string]string
	queryParams map[string]string
	body        string
	timeout     int // in milliseconds
}

// HttpResponse represents the response from an HTTP request
type HttpResponse struct {
	StatusCode int
	Status     string
	Headers    http.Header
	Body       string
}

// NewHttpRequest creates a new HttpRequest from the builder
func NewHttpRequest(b *HttpRequestBuilder) *HttpRequest {
	// Copy maps to ensure immutability
	headers := make(map[string]string)
	for k, v := range b.headers {
		headers[k] = v
	}

	queryParams := make(map[string]string)
	for k, v := range b.queryParams {
		queryParams[k] = v
	}

	return &HttpRequest{
		url:         b.url,
		method:      b.method,
		headers:     headers,
		queryParams: queryParams,
		body:        b.body,
		timeout:     b.timeout,
	}
}

// Execute performs the actual HTTP request and returns the response
func (r *HttpRequest) Execute() (*HttpResponse, error) {
	// Build the full URL with query parameters
	fullURL, err := r.buildURL()
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	// Create the request body
	var bodyReader io.Reader
	if r.body != "" {
		bodyReader = strings.NewReader(r.body)
	}

	// Create the HTTP request
	req, err := http.NewRequest(r.method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Duration(r.timeout) * time.Millisecond,
	}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return &HttpResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header,
		Body:       string(bodyBytes),
	}, nil
}

// buildURL constructs the full URL with query parameters
func (r *HttpRequest) buildURL() (string, error) {
	parsedURL, err := url.Parse(r.url)
	if err != nil {
		return "", err
	}

	// Add query parameters
	query := parsedURL.Query()
	for key, value := range r.queryParams {
		query.Set(key, value)
	}
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}

// Getters for HttpRequest fields
func (r *HttpRequest) URL() string                    { return r.url }
func (r *HttpRequest) Method() string                 { return r.method }
func (r *HttpRequest) Headers() map[string]string     { return r.headers }
func (r *HttpRequest) QueryParams() map[string]string { return r.queryParams }
func (r *HttpRequest) Body() string                   { return r.body }
func (r *HttpRequest) Timeout() int                   { return r.timeout }

func (r *HttpRequest) String() string {
	return fmt.Sprintf("HttpRequest{url='%s', method='%s', headers=%v, queryParams=%v, body='%s', timeout=%d}",
		r.url, r.method, r.headers, r.queryParams, r.body, r.timeout)
}

func (resp *HttpResponse) String() string {
	return fmt.Sprintf("HttpResponse{status='%s', statusCode=%d, bodyLength=%d}",
		resp.Status, resp.StatusCode, len(resp.Body))
}
