package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func createdByPostman() {

	url := "https://api3.useb.co.kr/ocr/idcard-driver"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("/Users/dong/Desktop/test-data/idcard/toUse.jpeg")
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("image", filepath.Base("/Users/dong/Desktop/test-data/idcard/toUse.jpeg"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE2NTg0NjM0MTQsImV4cCI6MTY1OTA2ODIxNCwiaXNzIjoiaHR0cHM6XC9cL2F1dGgudXNlYi5jby5rciIsImRhdGEiOnsiaWQiOiIzIiwidXNlcm5hbWUiOm51bGwsImVtYWlsIjoiYWJsZW1hbjgyQHVzZWIuY28ua3IiLCJwcmljZV9wbGFuIjpudWxsfSwic2NvcGVzIjpbImFkbWluIiwib2NyIiwic3RhdHVzIiwib2NyLWRvYyIsInN0YXR1cy1kb2MiLCJmYWNlIiwib3BlbmJhbmsiLCJtYXNraW5nIiwia2V5cyJdfQ.IjSPIeiT8ArPot1MkLmotosyiUIBKjI7lpVW4qRVwTk")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
