package main

import (
	"encoding/json"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/"

func fetchErrorHandled(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

func fetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(target); err != nil {
		return err
	}
	return nil
}
