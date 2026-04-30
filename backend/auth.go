package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const introspectPath = "/protocol/openid-connect/token/introspect"

type keycloakCfg struct {
	baseURL      string
	clientID     string
	clientSecret string
}

func keycloakConfig() keycloakCfg {
	return keycloakCfg{
		baseURL:      os.Getenv("KEYCLOAK_URL"),
		clientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		clientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	}
}

func auth(tokenString string) (*bool, error) {
	response, err := introspectToken(tokenString)
	if err != nil {
		return nil, err
	}
	// Check token validity
	if active, ok := response["active"].(bool); ok && active {
		fmt.Println("Token is valid!")
		fmt.Println("Token details:", response)

		found := false
		// Example: Retrieve roles
		if roles, ok := response["realm_access"].(map[string]interface{}); ok {
			if roleList, ok := roles["roles"].([]interface{}); ok {
				fmt.Println("Roles:", roleList)
				for _, role := range roleList {
					if role == "prothetic_user" {
						found = true
						break
					}
				}
			}
		}
		return &found, nil
	} else {
		return nil, fmt.Errorf("token is invalid or expired")
	}
}

func introspectToken(token string) (map[string]interface{}, error) {
	cfg := keycloakConfig()
	introspectionURL := fmt.Sprintf("%s%s", cfg.baseURL, introspectPath)

	data := url.Values{}
	data.Set("client_id", cfg.clientID)
	data.Set("client_secret", cfg.clientSecret)
	data.Set("token", token)

	fmt.Printf("query to keycloak url(%s) url_values(%s)\n", introspectionURL, data.Encode())
	// Create HTTP request
	req, err := http.NewRequest("POST", introspectionURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	fmt.Println("body from keycloak: ", string(body))
	// Parse JSON response
	var introspectionResponse map[string]interface{}
	if err := json.Unmarshal(body, &introspectionResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return introspectionResponse, nil
}
