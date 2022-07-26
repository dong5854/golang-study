package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OcrApi(c *gin.Context) {
	var requestBody ocrRequest
	bearerToken := getBearerToken()

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
