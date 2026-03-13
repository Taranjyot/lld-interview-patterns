package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept")

	if err != nil {
		fmt.Println("request creation error", err)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("request error:", err)
		return
	}

	defer resp.Body.Close()

	var post Post

	err = json.NewDecoder(resp.Body).Decode(&post)

	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fmt.Println("ID:", post.ID)
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
}
