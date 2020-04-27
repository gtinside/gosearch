package scrapper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Retrieve : Test for Retrieve function
func TestRetrieve(t *testing.T) {
	data, err :=ioutil.ReadFile("testdata/testAPIResult.json")
	if err != nil {
		panic(err)
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, string(data))
	}))
	defer testServer.Close()

	result := Retrieve("test", 3, testServer.URL)
	if len(result) != 3 {
		message := fmt.Sprintf("Invalid length: expected length 3, actual length %v", len(result))
		panic(message)
		t.Error()
	}

	if result[0] != "https://duckduckgo.com/Test_cricket" {
		message := fmt.Sprintf("Invalid URL in the response: expected %v, actual %v",
			"https://duckduckgo.com/Test_cricket", result[0])
		panic(message)
		t.Fail()
	}

	if result[1] != "https://duckduckgo.com/Test_(wrestler)" {
		message := fmt.Sprintf("Invalid URL in the response: expected %v, actual %v",
			"https://duckduckgo.com/Test_(wrestler)", result[1])
		panic(message)
		t.Fail()
	}
}
