package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com")
	if err != nil {
		// Deal with the error as you need:
		return nil, err
	}

	fmt.Printf("Status Code: %d\n", response.StatusCode)
	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Body: %s\n", response.String())

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		// Deal with the unmarshal error as you need:
		return nil, err
	}

	fmt.Printf("Repository URL: %s\n", endpoints.RepositoryUrl)
	return &endpoints, nil
}
