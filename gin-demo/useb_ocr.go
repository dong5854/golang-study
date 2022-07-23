package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

func usebOcr(c *gin.Context) {
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

	//saveErr := c.SaveUploadedFile(requestBody.File, filepath.Join("/Users/dong/test-files", requestBody.File.Filename))
	//
	//if saveErr != nil {
	//	println("저장 에러", saveErr.Error())
	//}

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

func usebIdcard(token string, requestBody *ocrRequest) ([]byte, error) {
	baseUrl := "https://api3.useb.co.kr/"
	bearer := "Bearer " + token
	var req *http.Request

	if requestBody.File != nil { // multipart-form 으로 받을 경우
		file, err := requestBody.File.Open()
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("image", requestBody.File.Filename)
		_, cpyErr := io.Copy(part, file)
		if cpyErr != nil {
			log.Fatalln(cpyErr)
		}
		writer.Close()

		fmt.Println("MULTIPART 요청", requestBody.File.Filename, writer.FormDataContentType())

		req, err = http.NewRequest("POST", baseUrl+"ocr/idcard-driver", body)
		req.Header.Add("Content-Type", writer.FormDataContentType())
	} else { // json 으로 받은 경우
		ocrRequestJSON, err := json.Marshal(requestBody)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("JSON 요청")
		req, err = http.NewRequest("POST", baseUrl+"ocr/idcard-driver", bytes.NewBuffer(ocrRequestJSON))
		req.Header.Add("Content-Type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}
	}

	req.Header.Add("Authorization", bearer)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	return data, err
}
