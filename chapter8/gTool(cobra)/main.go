package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/levigross/grequests"
	"github.com/spf13/cobra"
)

var (
	githubtoken    = os.Getenv("GITHUB_TOKEN")
	requestOptions = &grequests.RequestOptions{Auth: []string{githubtoken, "x-oauth-basic"}}
)

type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

type File struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		log.Fatalln("Uanable to make request: ", err)
	}
	return resp
}

func createGist(url string, args []string) *grequests.Response {
	//get first two arguments
	description := args[0]
	var fileContents = make(map[string]File)
	for i := 1; i < len(args); i++ {
		dat, err := os.ReadFile(args[i])
		if err != nil {
			log.Println("Please check the filename, Absolute Path (or) same directory are allowed")
			return nil
		}
		var file File
		file.Content = string(dat)
		fileContents[args[i]] = file
	}

	var gist = Gist{Description: description, Public: true, Files: fileContents}
	var postBody, _ = json.Marshal(gist)
	var requestOptions_copy = requestOptions
	//Add data to json fielf
	requestOptions_copy.JSON = string(postBody)
	//make post req to github
	resp, err := grequests.Post(url, requestOptions_copy)
	if err != nil {
		log.Println("Create request failed for Github APi")
	}
	return resp
}

func fetchCmd(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		var repos []Repo
		user := args[0]
		repoUrl := fmt.Sprintf("https://api.github.com/users/%s/repos", user)
		resp := getStats(repoUrl)
		resp.JSON(&repos)
		log.Println(repos)
	} else {
		log.Println("Please provide username. ")
	}
}

func createCmd(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		var postUrl = "https://api.github.com/gists"
		resp := createGist(postUrl, args)
		log.Println(resp.String())
	} else {
		log.Println("Please provide sufficient arguments, See -h for help")
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:     "githubAPItool",
		Version: "1.0",
	}
	fetchCmd := &cobra.Command{
		Use:     "fetch [user]",
		Aliases: []string{"f"},
		Short:   "Fetch the repo details belonging to this user",
		Example: "main.exe Fetch/f torvalds",
		Args:    cobra.ExactArgs(1),
		Run:     fetchCmd,
	}

	createCmd := &cobra.Command{
		Use:     "Create [name] [description] [files....]",
		Aliases: []string{"C", "c"},
		Short:   "Create github gists from given files/Texts",
		Example: "main.exe Create/ C/ c [file title] [file / file directiory ]",
		Args:    cobra.ExactArgs(2),
		Run:     createCmd,
	}

	rootCmd.AddCommand(fetchCmd, createCmd)
	rootCmd.Execute()
}
