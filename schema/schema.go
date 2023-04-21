package schema

import (
	"github.com/codebdy/entify"
	"github.com/codebdy/entify-graphql-schema/schema/parser"
	"github.com/graphql-go/graphql"
)

type MetaGraphqlSchema struct {
	QueryFields    []*graphql.Field
	MutationFields []*graphql.Field
	Directives     []*graphql.Directive
	Types          []graphql.Type
	proccessor     *MetaProcessor
}

type MetaProcessor struct {
	Repo        *entify.Repository
	modelParser parser.ModelParser
}

func New(r *entify.Repository) MetaGraphqlSchema {
	processor := &MetaProcessor{
		Repo: r,
	}

	processor.modelParser.ParseModel(r)
	return MetaGraphqlSchema{
		QueryFields:    processor.QueryFields(),
		MutationFields: processor.mutationFields(),
		Types:          processor.modelParser.EntityTypes(),
		proccessor:     processor,
	}
}

func (s *MetaGraphqlSchema) Parser() *parser.ModelParser {
	return &s.proccessor.modelParser
}

func (s *MetaGraphqlSchema) OutputType(name string) graphql.Type {
	return s.proccessor.modelParser.OutputType(name)
}
