package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func evalBlockStatement(block *ast.BlockStatement, enclosure *object.Environment) object.Object {
	var result object.Object
	var env = object.NewEnclosedEnvironment(enclosure)
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}
	return result
}
