package main

import (
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
	res, _ := fb.Get("/search", fb.Params{
		"access_token": "",
		"type":         "page",
		"q":            "nightlife,singapore",
	})

	var items []Page

	err := res.DecodeField("data", &items)
	if err != nil {
		log.Fatalln("An error has happened: ", err)
		return nil
	}

	return items
}
