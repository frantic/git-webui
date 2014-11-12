package main

import (
	"encoding/json"
	"fmt"
	"github.com/libgit2/git2go"
	"log"
	"net/http"
	"os"
)

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	file := path + "/" + r.URL.Path[1:]
	http.ServeFile(w, r, file)
}

type Commit struct {
	Sha1    string `json:"sha1"`
	Message string `json:"message"`
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	repo, err := git.OpenRepository(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	branches, _ := repo.NewBranchIterator(git.BranchLocal)
	branch, _, _ := branches.Next()
	fmt.Println(branch.Name())
	walk, _ := repo.Walk()
	commits := []Commit{}
	walk.PushHead()
	walk.Iterate(func(commit *git.Commit) bool {
		info := &Commit{
			Sha1:    commit.Id().String(),
			Message: commit.Message(),
		}
		commits = append(commits, *info)
		fmt.Print(info)
		return true
	})

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(commits)
	w.Write(b)
}

func main() {
	fmt.Println(os.Getwd())
	fmt.Println("Visit http://localhost:8080/")
	http.HandleFunc("/", frontendHandler)
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}
