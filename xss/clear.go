package xss

import (
	"fmt"
)

//ClearComment to report malicious comment on a given url
func (s *Client) ClearComment(num int) bool {
	status := false
	fullURL := fmt.Sprintf("%s/report/%d", s.URLBase, num)
	res, _ := s.getRequest(fullURL)

	if res.Result == "success" {
		status = true
	}
	return status
}
