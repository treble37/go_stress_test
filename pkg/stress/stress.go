package stress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"bytes"
)

type RequestBody struct {
	body []byte
}

type stressTest struct {
	url         string
	contentType string
	body        RequestBody
}

func NewRequestBody() *RequestBody {
	requestBody, _ := json.Marshal(map[string]string{
		"some": "data",
	})
	return &RequestBody{body: requestBody}
}
func NewDefaultStressTest() *stressTest {
	return &stressTest{url: "https://httpbin.org/anything", contentType: "application/json", body: *NewRequestBody()}
}
func NewStressTest(url string, contentType string, body RequestBody) *stressTest {
	return &stressTest{url: url, contentType: contentType, body: body}
}

// https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
// https://medium.com/dev-bits/making-concurrent-http-requests-in-go-programming-language-823b51bb1dc2
func (st *stressTest) PostRequest(ch chan<- string) {
	start := time.Now()
	resp, err := http.Post(st.url, st.contentType, bytes.NewBuffer(st.body.body))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	secs := time.Since(start).Seconds()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), st.url)
}
