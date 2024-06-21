package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	src := ""
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(src)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", feed)
	fmt.Println(feed.Title)
	for i, it := range feed.Items {
		fmt.Println("index", i+1)
		//fmt.Printf("%+v\n", it)
		fmt.Println("Title:", it.Title)
		fmt.Println("Link:", it.Link)
		fmt.Println("Published:", it.PublishedParsed)
		fmt.Println("Extensions:", it.Extensions["media"]["thumbnail"][0].Attrs["url"])
		fmt.Println()
	}
}
