package main

import (
	"context"
	"fmt"
	gql "gql/gen"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

func main() {
	const endpoint = "http://192.168.1.204:8088/graphql"
	gqlc := graphql.NewClient(endpoint, http.DefaultClient)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	versionRsp, err := gql.Version(ctx, gqlc)
	if err != nil {
		panic(err)
	}
	fmt.Println(versionRsp.GetVersion())
}
