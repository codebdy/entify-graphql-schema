package parser

import (
	"fmt"
)

var allSDL = `
#graphql-go-tools不支持V2
#extend schema
#	@link(url: "https://specs.apollo.dev/federation/v2.0", import: ["@key"])

scalar JSON
scalar DateTime

type Query {
%s
}

type Mutation {
%s
}
%s
`

func (p *ModelParser) MakeFederationSDL() string {
	sdl := allSDL
	queryFields, queryTypes := p.QuerySDL()
	mutationFields, mutationTypes := p.MutationSDL()
	return fmt.Sprintf(sdl, queryFields, mutationFields, queryTypes+mutationTypes)
}
