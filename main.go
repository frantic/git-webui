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

func getDiff() {
	path, _ := os.Getwd()
	repo, _ := git.OpenRepository(path)
	walk, _ := repo.Walk()
	var newTree, oldTree *git.Tree
	walk.PushHead()
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
	diff.ForEach(func(delta git.DiffDelta, i float64) (git.DiffForEachHunkCallback, error) {
		fmt.Println(delta.OldFile.Path, " -> ", delta.NewFile.Path)
		return func(hunk git.DiffHunk) (git.DiffForEachLineCallback, error) {
			fmt.Print(hunk.Header)
			return func(line git.DiffLine) error {
				if line.Origin == git.DiffLineAddition {
					fmt.Print("+", line.Content)
				} else {
					fmt.Print("-", line.Content)
				}
				return nil
			}, nil
		}, nil
	}, git.DiffDetailLines)
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
	// http.ListenAndServe(":8080", nil)
	getDiff()
}
