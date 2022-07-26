package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net/http"
	"time"
)

func getBearerToken() string {
	cfg, err := ini.Load("auth_info.ini")
	if err != nil {
		fmt.Println(err)
	}

	layout := "2006-01-02 15:04:05"
	expireDate := cfg.Section("API_TOKEN").Key("expires_in").String()
	currentDate := time.Now().Format(layout)

	if expireDate < currentDate {
		email := cfg.Section("AUTH").Key("email").String()
		password := cfg.Section("AUTH").Key("password").String()
		body := Auth{Email: email, Password: password}
		bodyBytes, marshalErr := json.Marshal(body)

		if marshalErr != nil {
			fmt.Println(marshalErr)
		}

		resp, respErr := http.Post("https://auth.useb.co.kr/oauth/token", "application/json", bytes.NewBuffer(bodyBytes))
		if respErr != nil {
			fmt.Println(respErr)
		}
		defer resp.Body.Close()

		data, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			fmt.Println(readErr)
		}

		var authResponse = AuthResponse{}
		unMarshalErr := json.Unmarshal(data, &authResponse)
		if unMarshalErr != nil {
			fmt.Println(unMarshalErr)
		}
		cfg.Section("API_TOKEN").Key("token").SetValue(authResponse.Jwt)
		cfg.Section("API_TOKEN").Key("expires_in").SetValue(authResponse.ExpiresIn)
		cfg.SaveTo("auth_info.ini")
	}
	bearerToken := cfg.Section("API_TOKEN").Key("token").String()
	return bearerToken
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	Jwt           string `json:"jwt"`
	ExpiresIn     string `json:"expires_in"`
	TransactionId string `json:"transaction_id"`
}
