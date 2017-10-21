package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type Result struct {
	Total int `json:"total_count"`
	Items []Item
}

type User struct {
	Login string
}

type Item struct {
	Url   string
	Title string
	User  User
	State string
}

func main() {
	resp, err := http.Get("https://api.github.com/search/issues?q=user:ipfs+user:libp2p&sort=updated")
	if err != nil {
		panic(err)
	}

	var res Result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		panic(err)
	}

	for i := 0; i < 20 && i < len(res.Items); i++ {
		item := res.Items[i]
		fmt.Printf("[%s] %s: %s\n", item.State, item.User.Login, color.BlueString("%q", item.Title))
		fmt.Printf("\t%s\n", item.Url)
	}
}
