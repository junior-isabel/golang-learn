package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CustomerWrite struct{}

type GithubRepository []struct {
	FullName string `json:"full_name"`
}

func (w CustomerWrite) Write(p []byte) (int, error) {

	var res GithubRepository
	err := json.Unmarshal(p, &res)
	if err != nil {
		return len(p), err
	}
	for _, r := range res {
		fmt.Println(r.FullName)
	}
	return len(p), err
}
func main() {

	res, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")

	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	write := CustomerWrite{}

	io.Copy(write, res.Body)
}
