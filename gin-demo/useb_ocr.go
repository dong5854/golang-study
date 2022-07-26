package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OcrApi(c *gin.Context) {
	var requestBody ocrRequest
	bearerToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE2NTg0NjM0MTQsImV4cCI6MTY1OTA2ODIxNCwiaXNzIjoiaHR0cHM6XC9cL2F1dGgudXNlYi5jby5rciIsImRhdGEiOnsiaWQiOiIzIiwidXNlcm5hbWUiOm51bGwsImVtYWlsIjoiYWJsZW1hbjgyQHVzZWIuY28ua3IiLCJwcmljZV9wbGFuIjpudWxsfSwic2NvcGVzIjpbImFkbWluIiwib2NyIiwic3RhdHVzIiwib2NyLWRvYyIsInN0YXR1cy1kb2MiLCJmYWNlIiwib3BlbmJhbmsiLCJtYXNraW5nIiwia2V5cyJdfQ.IjSPIeiT8ArPot1MkLmotosyiUIBKjI7lpVW4qRVwTk"

	if _, err := c.FormFile("image"); err == nil {
		err := c.Bind(&requestBody)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := c.BindJSON(&requestBody)
		if err != nil {
			fmt.Println(err)
		}
	}

	resp, err := usebIdcard(bearerToken, &requestBody)
	if err != nil {
		fmt.Println(err)
	}

	var responseObj = ocrResponse{}
	err = json.Unmarshal(resp, &responseObj)
	if err != nil {
		fmt.Println("Failed to json.Unmarshal", err)
	}

	c.JSON(http.StatusOK, responseObj)
}
