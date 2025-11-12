package emailgateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type priority string

const (
	PriorityLow  priority = "LOW"
	PriorityHigh priority = "HIGH"

	baseURL                      = "https://api.cm.com/email/gateway/v1"
	defaultHttpTimeout           = 30 * time.Second
	defaultTransactionalPriority = PriorityHigh
)

var (
	ErrUnableToSendEmail = errors.New("unable to send email")
)

type Client struct {
	baseClient            *http.Client
	productToken          string
	transactionalPriority priority
}

type Config struct {
	ProductToken                 string
	CustomHttpClient             *http.Client
	DefaultTransactionalPriority priority
}

func NewClient(config Config) (*Client, error) {
	var client Client

	if config.ProductToken == "" {
		return nil, errors.New("config.ProductToken must be set")
	}
	client.productToken = config.ProductToken

	client.baseClient = &http.Client{Timeout: defaultHttpTimeout}
	if config.CustomHttpClient != nil {
		client.baseClient = config.CustomHttpClient
	}

	client.transactionalPriority = defaultTransactionalPriority
	if config.DefaultTransactionalPriority != "" {
		client.transactionalPriority = config.DefaultTransactionalPriority
	}

	return &client, nil
}

type SendEmailResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	MessageID string `json:"messageId"`
}

func (c *Client) SendTransactionalEmail(email Email) (*SendEmailResponse, error) {
	emailPriority := c.transactionalPriority
	if email.Priority != "" {
		emailPriority = email.Priority
	}

	mailBody, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, baseURL+"/transactional?priority="+string(emailPriority), bytes.NewBuffer(mailBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CM-PRODUCTTOKEN", c.productToken)

	resp, err := c.baseClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response SendEmailResponse
	_ = json.NewDecoder(resp.Body).Decode(&response)

	if resp.StatusCode != http.StatusAccepted {
		return &response, ErrUnableToSendEmail
	}

	return &response, nil
}

func (c *Client) SendMarketingEmail(email Email) (*SendEmailResponse, error) {
	mailBody, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, baseURL+"/transactional", bytes.NewBuffer(mailBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CM-PRODUCTTOKEN", c.productToken)

	resp, err := c.baseClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response SendEmailResponse
	_ = json.NewDecoder(resp.Body).Decode(&response)

	if resp.StatusCode != http.StatusAccepted {
		return &response, ErrUnableToSendEmail
	}

	return &response, nil
}
