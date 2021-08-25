package codegen

import (
	"fmt"
	"go/ast"
	"go/token"
	"path"
	"reflect"
	"strings"
)

type TdElement struct {
	TypeName    string
	TypeKind    reflect.Kind
	PackageName string
	FileName    string
	Help        string
}

type TdConstant struct {
	TdElement
	ConstValue string
}

type TdPackage2 map[string]interface{}
type TdProto2 map[string]TdPackage2

func ParseTdproto2(infoToFill *TdProto2) error {

	allPackages := make(map[string]*ast.Package)
	err := extractTdprotoAst(&allPackages)
	if err != nil {
		return err
	}

	for packageName, packageAst := range allPackages {
		newPackageMap := make(TdPackage2)
		(*infoToFill)[packageName] = newPackageMap
		err := parseAst(packageName, &newPackageMap, packageAst)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseAst(packageName string, packageMap *TdPackage2, packageAst *ast.Package) error {
	for fileName, fileAst := range packageAst.Files {

		basePath := path.Base(fileName)
		basePathNoExt := strings.TrimRight(basePath, path.Ext(basePath))

		err := parseFileAst(packageMap, basePathNoExt, fileAst)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseFileAst(packageMap *TdPackage2, fileName string, fileAst *ast.File) error {
	for _, declaration := range fileAst.Decls {
		switch declarationType := declaration.(type) {
		case *ast.GenDecl:
			err := parseGenericDeclaration(packageMap, fileName, declarationType)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func parseGenericDeclaration(packageMap *TdPackage2, fileName string, genDeclaration *ast.GenDecl) error {
	switch genDeclaration.Tok {
	case token.CONST:
		return parseConstDeclaration2(packageMap, fileName, genDeclaration)
	case token.TYPE:
		for _, aSpec := range genDeclaration.Specs {
			aTypeSpec := aSpec.(*ast.TypeSpec)
			err := parseTypeDeclaration2(packageMap, genDeclaration, aTypeSpec, fileName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func parseConstDeclaration2(packageMap *TdPackage2, fileName string, genDeclaration *ast.GenDecl) error {
	for _, spec := range genDeclaration.Specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			return fmt.Errorf("expected const spec got %+v", spec)
		}

		if len(valueSpec.Names) != 1 {
			return fmt.Errorf("expected one constant name got %+v", valueSpec.Names)
		}

		constName := valueSpec.Names[0].Name

		constTypeName := fmt.Sprintf("%s", valueSpec.Type)
		if constTypeName == "" || valueSpec.Type == nil {
			errorLogger.Printf("WARN: const has no typeName %s", constName)
			continue
		}

		if len(valueSpec.Values) != 1 {
			return fmt.Errorf("expected one constant value got %+v", valueSpec.Values)
		}

		constValue, ok := valueSpec.Values[0].(*ast.BasicLit)
		if !ok {
			return fmt.Errorf("could not extract constant value %+v", valueSpec.Values[0])
		}

		(*packageMap)[constName] = TdConstant{
			ConstValue: constValue.Value,
			TdElement: TdElement{
				FileName: fileName,
				TypeKind: reflect.String,
				Help:     cleanHelp(valueSpec.Doc.Text()),
			},
		}

	}

	return nil
}

func parseTypeDeclaration2(packageMap *TdPackage2, genDeclaration *ast.GenDecl, declarationSpec *ast.TypeSpec, fileName string) error {

	return nil
}
