package xss

import (
	"fmt"
	"log"
	"net/http"
)

//GetSearchPage will get request to user comment 2
func (s *Client) GetSearchPage() string {
	fullURL := fmt.Sprintf("%s/user/3", s.URLBase)
	res, err := http.Get(fullURL)
	if err != nil {
		log.Fatal("Cant send get request ", err)
	}
	defer res.Body.Close()
	body := redBody(res)
	base := parseHTML(string(body))
	s.ClearComment(3)
	return decodePayload(base)
}
