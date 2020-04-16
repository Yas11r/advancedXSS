package xss

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	payload string
)

//Payload extract
type Payload struct {
	EncodedPayload string
}

//RawPayload from xss.js file
func RawPayload() []byte {
	data, err := ioutil.ReadFile("./js/xss.js")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	return data
}

//redBody read the body html
func redBody(resp *http.Response) []byte {

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	return data
}

//EncodePayload payload with Base64 URL Encoded
func EncodePayload(payload []byte) string {
	return b64.URLEncoding.EncodeToString(payload)
}

//DecodePayload Decodes payload
func decodeURLPayload(payload string) string {
	un, _ := b64.URLEncoding.DecodeString(payload)
	return string(un)
}

//DecodePayload Decodes payload
func decodePayload(payload string) string {
	un, _ := b64.StdEncoding.DecodeString(payload)
	return string(un)
}

//GetPayload export it
func (s *Client) GetPayload() *Payload {
	payload := EncodePayload(RawPayload())
	return &Payload{payload}
}
