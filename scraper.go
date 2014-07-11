package main

import (
	"flag"
	fb "github.com/huandu/facebook"
	"log"
)

func main() {
	items := getPages()

	for _, item := range items {
		log.Println(item.Id)
	}
}

type Page struct {
	Id           string
	Category     string
	CategoryList []Category
	Name         string
}

type Category struct {
	Id   string
	Name string
}

func getPages() []Page {
	// fb access token
	var accessToken = flag.String("accessToken", "", "fb access token")

	// fb element e.g. page, place
	var fbType = flag.String("fbType", "", "fb element")

	// query e.g. night,philippines
	var query = flag.String("query", "", "query")

	flag.Parse()

	log.Printf("fbType=%s, query=%q", *fbType, *query)

	res, _ := fb.Get("/search", fb.Params{
		"access_token": *accessToken,
		"type":         *fbType,
		"q":            *query,
	})

	var items []Page

	err := res.DecodeField("data", &items)
	if err != nil {
		log.Fatalln("An error has happened: ", err)
		return nil
	}

	return items
}
