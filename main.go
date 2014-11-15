package main

import (
	"encoding/json"
	"fmt"
	"github.com/libgit2/git2go"
	"log"
	"net/http"
	"os"
	"strings"
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

func diffHandler(w http.ResponseWriter, r *http.Request) {
	sha := strings.Split(r.URL.Path, "/")[2]
	path, _ := os.Getwd()
	repo, _ := git.OpenRepository(path)
	walk, _ := repo.Walk()
	var newTree, oldTree *git.Tree
	oid, x := git.NewOid(sha)
	if x != nil {
		log.Fatal(x)
	}
	walk.Push(oid)
	walk.Iterate(func(commit *git.Commit) bool {
		if newTree == nil {
			newTree, _ = commit.Tree()
			return true
		}
		oldTree, _ = commit.Tree()
		return false
	})
	opts := &git.DiffOptions{}
	diff, err := repo.DiffTreeToTree(oldTree, newTree, opts)
	if err != nil {
		log.Fatal(err)
	}
	s := ""
	diff.ForEach(func(delta git.DiffDelta, i float64) (git.DiffForEachHunkCallback, error) {
		s = s + delta.OldFile.Path + " -> " + delta.NewFile.Path + "\n"
		return func(hunk git.DiffHunk) (git.DiffForEachLineCallback, error) {
			s = s + hunk.Header
			return func(line git.DiffLine) error {
				if line.Origin == git.DiffLineAddition {
					s = s + "+" + line.Content
				} else {
					s = s + "-" + line.Content
				}
				return nil
			}, nil
		}, nil
	}, git.DiffDetailLines)
	b, _ := json.Marshal(s)
	w.Write(b)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	repo, err := git.OpenRepository(path)
	if err != nil {
		log.Fatal(err)
	}
	walk, _ := repo.Walk()
	commits := []Commit{}
	walk.PushHead()
	walk.Iterate(func(commit *git.Commit) bool {
		info := &Commit{
			Sha1:    commit.Id().String(),
			Message: commit.Message(),
		}
		commits = append(commits, *info)
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
	http.HandleFunc("/diff/", diffHandler)
	http.ListenAndServe(":8080", nil)
}
