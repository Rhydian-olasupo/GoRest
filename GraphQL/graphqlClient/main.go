package main

import (
	"github.com/machinebox/graphql"
)

//Response of API

type Response struct {
	License struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"license"`
}

func main() {
	//create a cliennt (safe to share across requests)

	client := graphql.NewClient("https://api.github.com/graphql")

	//make a request to Github API

	req := graphql.NewRequest(`
		query {
			license(key: "apache-2.0") {
				name
				description
			}
		}
	`)

}
