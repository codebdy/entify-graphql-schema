package resolve

import (
	"github.com/codebdy/entify-core/shared"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func parseListFields(info graphql.ResolveInfo) []string {
	fields := []string{}
	if len(info.FieldASTs) > 0 {
		nodesSelections := info.FieldASTs[0].SelectionSet.Selections
		for _, selection := range nodesSelections {
			noesField, ok := selection.(*ast.Field)
			if ok && noesField.Name.Value == shared.NODES {
				for _, fieldSelection := range noesField.SelectionSet.Selections {
					field, ok := fieldSelection.(*ast.Field)
					if ok {
						fields = append(fields, field.Name.Value)
					}
				}
			}
		}
	}
	return fields
}
