package resolve

import (
	"context"

	"github.com/codebdy/entify"
	"github.com/codebdy/entify/model/meta"
	"github.com/codebdy/entify/shared"
	"github.com/codebdy/minions-go"
	"github.com/codebdy/minions-go/dsl"
	"github.com/graphql-go/graphql"
)

func LogicFlowMethodResolveFn(method *meta.MethodMeta, repository *entify.Repository, subFlows *[]dsl.SubLogicFlowMeta) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		var returnValue any
		ctx := minions.AttachSubFlowsToContext(subFlows, context.Background())
		logicFlow := minions.NewLogicflow(method.LogicMetas, ctx)

		output := logicFlow.Jointers.GetSingleOutput()

		if output != nil {
			output.Connect(func(inputValue any, ctx context.Context) {
				returnValue = inputValue
			})
		}

		input := logicFlow.Jointers.GetSingleInput()
		if input == nil {
			panic("No input")
		}
		//后端args是map，前端是数组。后端无法保证参数顺序，前端无法获知参数名字
		input.Push(p.Args, ctx)
		return returnValue, nil
	}
}
