package graphql

import (
	"context"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
	"os"
)

var QueryRelease struct {
	Repository struct {
		NameWithOwner graphql.String
	} `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`
}

func DoReleaseQuery(repositoryOwner string, repositoryName string) {
	var variables = map[string]interface{}{
		"repositoryOwner": graphql.String(repositoryOwner),
		"repositoryName":  graphql.String(repositoryName),
	}
	query := QueryRelease
	doQuery(query, variables)

}

func doQuery(query interface{}, variables map[string]interface{}) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient("https://api.github.com/graphql", httpClient, nil)
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		panic(err)
	}
}
