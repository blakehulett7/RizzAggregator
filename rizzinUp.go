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
	return rss
}

func ProcessRizz(rizzStruct Rss) {
	postArray := []post{}
	for _, post := range rizzStruct.Channel.Items {
		postArray = append(postArray, post)
	}
	for _, post := range rizzStruct.Entries {
		postArray = append(postArray, post)
	}
	for _, post := range postArray {
		fmt.Println(post.GetTitle())
	}
}
