package codegen

import (
	"fmt"
	"go/ast"
	"go/token"
	"path"
	"strings"
)

const StringTypeName = "string"

type TdElement struct {
	PackageName string
	FileName    string
	Help        string
}

type TdConstant struct {
	TdElement
	ConstValue string
	TypeName   string
}

type TdType2 struct {
	TdElement
	TypeName string
}

type TdTypeArray struct {
	TdElement
	TypeName string
}

type TdTypeMap struct {
	TdElement
	KeyTypeStr       string
	KeyTypePackage   string
	ValueTypeStr     string
	ValueTypePackage string
}

type TdPackage2 map[string]interface{}
type TdProto2 map[string]TdPackage2

type ParserState struct {
	CurrentFile        string
	CurrentPackageName string
	CurrentPackageMap  *TdPackage2
}

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

		err := parseFileAst(ParserState{
			CurrentFile:        basePathNoExt,
			CurrentPackageName: packageName,
			CurrentPackageMap:  packageMap,
		}, fileAst)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseFileAst(parser ParserState, fileAst *ast.File) error {
	for _, declaration := range fileAst.Decls {
		switch declarationType := declaration.(type) {
		case *ast.GenDecl:
			err := parseGenericDeclaration(parser, declarationType)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func parseGenericDeclaration(parser ParserState, genDeclaration *ast.GenDecl) error {
	switch genDeclaration.Tok {
	case token.CONST:
		return parseConstDeclaration2(parser, genDeclaration)
	case token.TYPE:
		for _, aSpec := range genDeclaration.Specs {
			aTypeSpec := aSpec.(*ast.TypeSpec)
			err := parseTypeDeclaration2(parser, genDeclaration, aTypeSpec)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func parseConstDeclaration2(parser ParserState, genDeclaration *ast.GenDecl) error {
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
			constTypeName = StringTypeName
		}

		if len(valueSpec.Values) != 1 {
			return fmt.Errorf("expected one constant value got %+v", valueSpec.Values)
		}

		var constValue string

		switch constExpr := valueSpec.Values[0]; constExpr.(type) {
		case *ast.BasicLit:
			constValue = constExpr.(*ast.BasicLit).Value
		case *ast.CallExpr:
			callExpr := constExpr.(*ast.CallExpr)
			constTypeName = callExpr.Fun.(*ast.Ident).Name
			constValue = callExpr.Fun.(*ast.Ident).Obj.Kind.String()
		}

		if !ok {
			return fmt.Errorf("could not extract constant value %+v", valueSpec.Values[0])
		}

		(*parser.CurrentPackageMap)[constName] = TdConstant{
			TdElement: TdElement{
				PackageName: parser.CurrentPackageName,
				FileName:    parser.CurrentFile,
				Help:        cleanHelp(valueSpec.Doc.Text()),
			},
			ConstValue: constValue,
			TypeName:   constTypeName,
		}

	}

	return nil
}

func parseTypeDeclaration2(parser ParserState, genDeclaration *ast.GenDecl, declarationSpec *ast.TypeSpec) error {

	helpString := cleanHelp(genDeclaration.Doc.Text())

	switch typeAst := declarationSpec.Type.(type) {
	case *ast.Ident:
		err := parseTypeDefinition2(parser, declarationSpec, typeAst, helpString)
		if err != nil {
			return err
		}
	case *ast.StructType:

	case *ast.ArrayType:
		err := parseArrayDefinition2(parser, declarationSpec, typeAst, helpString)
		if err != nil {
			return err
		}
	case *ast.MapType:
		err := parseMapTypeDefinition(parser, declarationSpec, typeAst, helpString)
		if err != nil {
			return nil
		}
	default:
		errorLogger.Printf("WARN: Not implemented type declaration %#v in file %s", typeAst, parser.CurrentFile)
	}

	return nil
}

func parseTypeDefinition2(parser ParserState, declarationSpec *ast.TypeSpec, typeIndent *ast.Ident, helpString string) error {

	typeName := declarationSpec.Name.Name

	(*parser.CurrentPackageMap)[typeName] = TdType2{
		TdElement: TdElement{

			Help:        helpString,
			FileName:    parser.CurrentFile,
			PackageName: parser.CurrentPackageName,
		},
		TypeName: typeIndent.Name,
	}

	return nil
}

func parseArrayDefinition2(parser ParserState, declarationSpec *ast.TypeSpec, arrayAst *ast.ArrayType, helpString string) error {
	typeName := declarationSpec.Name.Name
	arrayExpressionAst := arrayAst.Elt.(*ast.Ident)
	arrayTypeStr := arrayExpressionAst.Name

	(*parser.CurrentPackageMap)[typeName] = TdType2{
		TdElement: TdElement{

			Help:        helpString,
			FileName:    parser.CurrentFile,
			PackageName: parser.CurrentPackageName,
		},
		TypeName: arrayTypeStr,
	}

	return nil
}

func parseMapTypeDefinition(parser ParserState, declarationSpec *ast.TypeSpec, mapAst *ast.MapType, helpString string) error {
	typeName := declarationSpec.Name.Name
	var keyPackageName, keyTypeStr string

	err := parseExprToString(mapAst.Key, &keyPackageName, &keyTypeStr)
	if err != nil {
		return err
	}

	var valuePackageName, valueTypeStr string
	err = parseExprToString(mapAst.Value, &valuePackageName, &valueTypeStr)
	if err != nil {
		return err
	}

	(*parser.CurrentPackageMap)[typeName] = TdTypeMap{
		TdElement: TdElement{

			Help:        helpString,
			FileName:    parser.CurrentFile,
			PackageName: parser.CurrentPackageName,
		},
		KeyTypeStr:       keyTypeStr,
		KeyTypePackage:   keyPackageName,
		ValueTypeStr:     valueTypeStr,
		ValueTypePackage: valuePackageName,
	}

	return nil
}
