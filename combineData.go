package main

import (
	"encoding/json"
	"net/http"
)

type CombineResult struct {
	PostId        int    `json:"postId"`
	PostName      string `json:"postName"`
	CommentsCount int    `json:"commentsCount"`
	UserName      string `json:"userName"`
	Body          string `json:"body"`
}

func combineData(w http.ResponseWriter, r *http.Request) {
	var combinedResults []CombineResult

	comments, err := fetchComments()
	if err != nil {
		fetchErrorHandled(w, err)
	}

	posts, err := fetchPosts()
	if err != nil {
		fetchErrorHandled(w, err)
	}

	users, err := fetchUsers()
	if err != nil {
		fetchErrorHandled(w, err)
	}
	for _, post := range posts {
		var commentsCount int
		for _, comment := range comments {
			if comment.PostId == post.Id {
				commentsCount++
			}
		}

		var userName string
		for _, user := range users {
			if user.Id == post.UserId {
				userName = user.Username
				break
			}
		}

		combinedResult := CombineResult{post.Id, post.Title, commentsCount, userName, post.Body}

		combinedResults = append(combinedResults, combinedResult)
	}
	resultJSON, err := json.Marshal(combinedResults)
	if err != nil {
		fetchErrorHandled(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resultJSON); err != nil {
		fetchErrorHandled(w, err)
	}
}
