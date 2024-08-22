package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func FetchFeed(url string) Rss {
	responseData, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return Rss{}
	}
	body, err := io.ReadAll(responseData.Body)
	if err != nil {
		fmt.Println(err)
		return Rss{}
	}
	rss := Rss{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println(err)
		return Rss{}
	}
	fmt.Println(rss)
	return rss
}

func ProcessRizz(rizzStruct Rss) {
	titleArray := []string{}
	for _, post := range rizzStruct.Channel.Item {
		titleArray = append(titleArray, post.Title)
	}
	for _, post := range rizzStruct.Entry {
		titleArray = append(titleArray, post.Title)
	}
	for _, post := range titleArray {
		fmt.Println(post)
	}
}
