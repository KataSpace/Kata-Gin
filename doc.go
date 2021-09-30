package KataGin

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// Get the name and path of a func
func funcPathAndName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// Get the name of a func (with package path)
func funcName(f interface{}) string {
	splitFuncName := strings.Split(funcPathAndName(f), ".")
	return splitFuncName[len(splitFuncName)-1]
}

// Get description of a func
func funcDescription(f interface{}, name string) string {
	fileName, _ := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).FileLine(0)
	funcName := funcName(f)
	fset := token.NewFileSet()

	// Parse src
	parsedAst, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst

	importPath, _ := filepath.Abs("/")

	myDoc := doc.New(pkg, importPath, doc.AllDecls)

	for _, theStruct := range myDoc.Types{
		if theStruct.Name == name{
			for _, methodOfStruct := range theStruct.Methods{
				if methodOfStruct.Name == funcName{
					return strings.TrimSpace(methodOfStruct.Doc)
				}
			}
		}
	}

	for _, theFunc := range myDoc.Funcs {
		if theFunc.Name == funcName {
			return strings.TrimSpace(theFunc.Doc)
		}
	}

	return ""
}
