package resolve

import (
	"fmt"

	"github.com/codebdy/entify"
	"github.com/codebdy/entify-graphql-schema/resolve/script"
	"github.com/codebdy/entify/model/meta"
	"github.com/codebdy/entify/shared"
	"github.com/dop251/goja"
	"github.com/graphql-go/graphql"
)

func ScriptMethodResolveFn(code string, methodArgs []meta.ArgMeta, repository *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		scriptService := script.NewService(p.Context, repository)
		vm := goja.New()
		script.Enable(vm)

		var meMap map[string]interface{}

		vm.Set("$args", p.Args)
		vm.Set("$beginTx", scriptService.BeginTx)
		vm.Set("$clearTx", scriptService.ClearTx)
		vm.Set("$commit", scriptService.Commit)
		vm.Set("$rollback", scriptService.Rollback)
		vm.Set("$save", scriptService.Save)
		vm.Set("$saveOne", scriptService.SaveOne)
		vm.Set("$log", scriptService.WriteLog)
		vm.Set("$notice", scriptService.EmitNotification)
		//vm.Set("$query", scriptService.Query)
		vm.Set("$me", meMap)

		script.Enable(vm)
		funcStr := fmt.Sprintf(
			`
			%s
			`,
			code,
		)

		result, err := vm.RunString(funcStr)
		if err != nil {
			panic(err)
		}
		return result.Export(), nil
	}
}
