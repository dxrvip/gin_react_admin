package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type SystenmMenu struct {
	Package string         `json:"package"`
	Func    []FunctionInfo `json:"func"`
}

// 菜单生成。
type FunctionInfo struct {
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
}

// 解析go 源文件并提取函数信息生成菜单
func ParseApiFiles(dir string) ([]SystenmMenu, error) {
	var funcs []SystenmMenu
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), ".go") {
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				return err
			}
			if node.Doc != nil {
				docText := node.Doc.List[0].Text
				// 提取注释中的内容
				if strings.HasPrefix(docText, "/*") {

					content := strings.TrimSpace(strings.TrimPrefix(node.Doc.List[0].Text, "/*"))
					content = strings.TrimSpace(strings.TrimSuffix(content, "*/"))
					funcs = append(funcs, SystenmMenu{Package: content})
				}
				// 遍历文件中的声明
				for _, decl := range node.Decls {
					if f, ok := decl.(*ast.FuncDecl); ok {
						alias, description := getFunctionAliasAndDesctiption(f)
						if alias != "" || description != "" {
							funcs[len(funcs)-1].Func = append(funcs[len(funcs)-1].Func, FunctionInfo{
								Name:        f.Name.Name,
								Alias:       alias,
								Description: description,
							})
						}

					}
				}
			}

		}
		return nil

	})
	return funcs, err
}

// 从函数注释中获取中文别名
func getFunctionAliasAndDesctiption(f *ast.FuncDecl) (string, string) {
	var (
		alias       string
		description string
	)
	if f.Doc != nil {

		for _, comment := range f.Doc.List {
			// 假设使用 //@Summary: 中文别名
			if strings.HasPrefix(comment.Text, "// @Summary") {
				alias = strings.TrimSpace(comment.Text[len("// @Summary"):])
			} else if strings.HasPrefix(comment.Text, "// @Description:") {
				description = strings.TrimSpace(comment.Text[len("// @Description:"):])
			}
		}
	}
	return alias, description
}
