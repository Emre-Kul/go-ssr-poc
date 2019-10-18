package services

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	// "math/rand"
)

type News struct {
	Text string
	Type string
}

func FetchNews() News {
	var url string = "http://cat-fact.herokuapp.com/facts/591d9bb2227c1a0020d26826"
	var news News
	resp, err := http.Get(url)
	if(err != nil) {
		fmt.Println(err)
		return News{"a","b"}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if(err != nil) {
		fmt.Println(err)
		return News{"a","b"}
	}
	
	json.Unmarshal([]byte(body), &news)
	fmt.Println(news.Text);
    return news
}

func FetchNews2() string {
	// return "{ \"a\": \"" + string(rand.Int()) + "\"}"
	// return "{}"
	var url string = "http://cat-fact.herokuapp.com/facts/591d9bb2227c1a0020d26826"
	resp, err := http.Get(url)
	if(err != nil) {
		fmt.Println(err)
		return "{}"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if(err != nil) {
		fmt.Println(err)
		return "{}"
	}
	
	
    return string(body)
}