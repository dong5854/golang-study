package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		imgBase64Str := base64.StdEncoding.EncodeToString(fileData)
		requestBody.ImageBase64 = imgBase64Str
	}

	ocrRequestJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest("POST", baseUrl+"ocr/idcard-driver", bytes.NewBuffer(ocrRequestJSON))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
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
