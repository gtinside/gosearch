package scrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Topics : Struct for the result JSON
type Topics struct {
	FirstURL string
}

// Result : Struct for defining the result
type Result struct {
	RelatedTopics []Topics
}

// Retrieve : To query DuckDuckGo API and return back the results
func Retrieve(keyword string, num int, domain string) []string {
	output := make([]string, num)
	url := fmt.Sprintf("%s?q=%s&format=json", domain, keyword)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error retrieving results %s\n\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var result Result
		json.Unmarshal([]byte(string(data)), &result)
		for i, s := range result.RelatedTopics {
			if i == num-1 {
				break
			}
			output[i] = s.FirstURL
		}
	}
	return output

}