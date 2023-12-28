package main

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func fetchPosts() (posts []Post, err error) {
	err = fetchData(URL+"posts", &posts)
	return
}
