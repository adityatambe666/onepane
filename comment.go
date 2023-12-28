package main

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func fetchComments() (comments []Comment, err error) {
	err = fetchData(URL+"comments", &comments)
	return
}
