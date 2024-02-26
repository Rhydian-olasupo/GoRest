package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/levigross/grequests"
	"github.com/urfave/cli/v2"
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

func main() {
	app := cli.NewApp()
	// define command for our client
	app.Commands = []*cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "Fetch the repo details with user. [Usage] : githubAPI fetch user",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() > 0 {
					//Github API Logic
					var repos Repo
					user := ctx.Args().Get(0)
					var repoUrl = fmt.Sprintf("https://api.github.com/users/%s/repos", user)
					resp := getStats(repoUrl)
					resp.JSON(&repos)
					log.Println(repos)
				} else {
					log.Println("Please give me a username. See -h to see help")
				}
				return nil
			},
		},

		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Creates a gist from the given text. [Usage]:githubAPI name 'description' sample.txt",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() > 1 {
					//Github API logic
					args := ctx.Args()
					var postUrl = "https://api.github.com/gists"
					resp := createGist(postUrl, args.Slice())
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments See -h to see help")
				}
				return nil
			},
		},
	}

	app.Version = "1.0"
	app.Run(os.Args)

}
