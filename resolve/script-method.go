package resolve

import (
	"fmt"
	"strings"

	"github.com/codebdy/entify-core"
	"github.com/codebdy/entify-core/model/meta"
	"github.com/codebdy/entify-core/shared"
	"github.com/codebdy/entify-graphql-schema/resolve/script"
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

		var codes string
		for i := range repository.Model.Meta.Codes {
			codes = codes + repository.Model.Meta.Codes[i].ScriptText + "\n"
		}

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
			%s
			function %s(){
				%s
			%s
			}
			`,
			codes,
			method.Name,
			argsStr,
			method.LogicScript,
		)
		_, err := vm.RunString(funcStr)
		if err != nil {
			panic(err)
		}
		var fn func() interface{}
		err = vm.ExportTo(vm.Get(method.Name), &fn)
		if err != nil {
			panic(err)
		}

		return fn(), nil
	}
}
