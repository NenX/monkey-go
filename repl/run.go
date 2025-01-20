package repl

import (
	"io/ioutil"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

func RunFile(filePath string) error {
	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 创建环境
	env := object.NewEnvironment()

	// 创建lexer和parser
	l := lexer.New(string(content))
	p := parser.New(l)

	// 解析程序
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		// 如果有解析错误，打印错误
		for _, msg := range p.Errors() {
			println("parser error:", msg)
		}
		return nil
	}

	// 执行程序
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		println("evaluated:", evaluated.Inspect())
	}

	return nil
}
