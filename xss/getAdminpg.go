package xss

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//GetAdminPage will get request to user comment 2
func (s *Client) GetAdminPage() string {
	fullURL := fmt.Sprintf("%s/user/2", s.URLBase)
	res, err := http.Get(fullURL)
	if err != nil {
		log.Fatal("Cant send get request ", err)
	}
	defer res.Body.Close()

	body := redBody(res)
	base := parseHTML(string(body))
	//s.ClearComment(2)

	return strings.Split(decodePayload(base), "\n")[0]
}
