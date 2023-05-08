package schema

import (
	"github.com/codebdy/entify-graphql-schema/consts"
	"github.com/codebdy/entify/shared"
	"github.com/graphql-go/graphql"
)

func (m *MetaProcessor) serviceField() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(
			graphql.ObjectConfig{
				Name: consts.SERVICE_TYPE,
				Fields: graphql.Fields{
					consts.SDL: &graphql.Field{
						Type:        graphql.String,
						Description: "Service SDL",
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							defer shared.PrintErrorStack()
							return m.modelParser.MakeFederationSDL(), nil
						},
					},
				},
				Description: "_Service type of federation schema specification, and extends other fields",
			},
		),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			defer shared.PrintErrorStack()
			return map[string]interface{}{}, nil
		},
	}
}
