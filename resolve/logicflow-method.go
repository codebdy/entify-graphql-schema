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
		var returnValue any
		ctx := runtime.AttachSubFlowsToContext(subFlows, context.Background())
		logicFlow := runtime.NewLogicflow(method.LogicMeta, ctx)

		logicFlow.Jointers.GetSingleOutput().Connect(func(inputValue any, ctx context.Context) {
			returnValue = inputValue
		})

		//后端args是map，前端是数组。后端无法保证参数顺序，前端无法获知参数名字
		logicFlow.Jointers.GetSingleInput().Push(p.Args, ctx)
		return returnValue, nil
	}
}
