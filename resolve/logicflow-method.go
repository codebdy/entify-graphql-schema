package resolve

import (
	"context"

	"github.com/codebdy/entify"
	"github.com/codebdy/entify/model/meta"
	"github.com/codebdy/entify/shared"
	"github.com/codebdy/minions-go/dsl"
	"github.com/codebdy/minions-go/runtime"
	"github.com/graphql-go/graphql"
)

func LogicFlowMethodResolveFn(method *meta.MethodMeta, repository *entify.Repository, subFlows *[]dsl.SubLogicFlowMeta) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()

		logicFlow := runtime.NewLogicflow(method.LogicMeta, runtime.AttachSubFlowsToContext(subFlows, context.Background()) )

		return nil, nil
	}
}
