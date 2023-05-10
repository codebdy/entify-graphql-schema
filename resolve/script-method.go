package resolve

import (
	"fmt"
	"strings"

	"github.com/codebdy/entify"
	"github.com/codebdy/entify-graphql-schema/resolve/script"
	"github.com/codebdy/entify/model/meta"
	"github.com/codebdy/entify/shared"
	"github.com/dop251/goja"
	"github.com/graphql-go/graphql"
)

func ScriptMethodResolveFn(method *meta.MethodMeta, repository *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		entifyService := script.NewService(p.Context, repository)
		vm := goja.New()
		script.Enable(vm)

		var meMap map[string]interface{}

		vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
		vm.Set("$args", p.Args)
		vm.Set("$entify", entifyService)
		//vm.Set("$query", scriptService.Query)
		vm.Set("$me", meMap)

		argsStr := ""
		if len(p.Args) > 0 {
			templateStr := `const {%s} = $args`

			keys := make([]string, 0, len(method.Args))
			for _, arg := range method.Args {
				keys = append(keys, arg.Name)
			}

			argsStr = fmt.Sprintf(templateStr, strings.Join(keys, ","))

		}

		funcStr := fmt.Sprintf(
			`
			function %s(){
				%s
			%s
			}
			`,
			method.Name,
			argsStr,
			method.LogicScript,
		)
		_, err := vm.RunString(funcStr)
		if err != nil {
			panic(err)
		}
		var fn func() string
		err = vm.ExportTo(vm.Get(method.Name), &fn)
		if err != nil {
			panic(err)
		}

		return fn(), nil
	}
}
