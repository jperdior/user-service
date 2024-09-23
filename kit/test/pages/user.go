package pages

import "net/http"

const BaseUrl = "/users/"

func GetUser(id string, token *string) (*http.Request, error) {
	req, err := http.NewRequest("GET", BaseUrl+id, nil)

	if err != nil {
		return nil, err
	}
	return req, nil
}
