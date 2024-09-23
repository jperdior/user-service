package pages

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const BaseUrl = "/users"

type UserPage struct {
	JwtToken *string
}

func NewUserPage(jwtToken *string) *UserPage {
	return &UserPage{JwtToken: jwtToken}
}

func (up *UserPage) GetUser(id string) (*http.Request, error) {
	req, err := http.NewRequest("GET", BaseUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	if up.JwtToken != nil {
		req.Header.Set("Authorization", "Bearer "+*up.JwtToken)
	}
	return req, nil
}

func (up *UserPage) GetUsers() (*http.Request, error) {
	req, err := http.NewRequest("GET", BaseUrl, nil)
	if err != nil {
		return nil, err
	}
	if up.JwtToken != nil {
		req.Header.Set("Authorization", "Bearer "+*up.JwtToken)
	}
	return req, nil
}

func (up *UserPage) UpdateUser(id string, payload map[string]string) (*http.Request, error) {
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("PUT", BaseUrl+"/"+id, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if up.JwtToken != nil {
		req.Header.Set("Authorization", "Bearer "+*up.JwtToken)
	}
	return req, nil
}

func (up *UserPage) RegisterUser(payload map[string]string) (*http.Request, error) {
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
