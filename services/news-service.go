package services

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
)

type News struct {
	Text string
	Type string
}

func FetchNews() News {
	var url string = "http://cat-fact.herokuapp.com/facts/591d9bb2227c1a0020d26826"
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	var news News
	json.Unmarshal([]byte(body), &news)
	fmt.Println(news.Text);
    return news
}

func FetchNews2() []byte {
	var url string = "http://cat-fact.herokuapp.com/facts/591d9bb2227c1a0020d26826"
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
    return body
}