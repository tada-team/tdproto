package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/tada-team/tdproto"
)

var GolangPrimitiveTypes = map[string]string{
	"string":            "",
	"int":               "",
	"int64":             "",
	"uint16":            "",
	"uint":              "",
	"bool":              "",
	"interface{}":       "",
	"ISODateTimeString": "",
	"time.Time":         "",
}

type TdConstFields struct {
	Name  string
	Type  string
	Value string
	Help  string
}

type TdStructField struct {
	Name            string
	Help            string
	JsonName        string
	TypeStr         string
	IsPrimitive     bool
	IsReadOnly      bool
	IsPointer       bool
	IsList          bool
	IsOmitEmpty     bool
	IsNotSerialized bool
}

type TdStruct struct {
	Name             string
	Help             string
	Fields           []TdStructField
	ReadOnly         bool
	AnonnymousFields []string
	FileName         string
}

type TdType struct {
	Name     string
	Help     string
	IsArray  bool
	BaseType string
}

type TdInfo struct {
	TdStructs map[string]TdStruct
	TdTypes   []TdType
	TdEvents  map[string]string
	TdConsts  []TdConstFields
}

type TdEnum struct {
	Name   string
	Values []string
}

func (i TdInfo) GetEnums() []TdEnum {
	constMap := make(map[string][]string)

	for _, aConst := range i.TdConsts {
		constType := aConst.Type
		constValue := aConst.Value

		constValueList := constMap[constType]
		constMap[constType] = append(constValueList, strings.Trim(constValue, `"`))
	}

	var listOfEnums []TdEnum

	for key, value := range constMap {
		listOfEnums = append(listOfEnums, TdEnum{
			Name:   key,
			Values: value,
		})
	}

	return listOfEnums
}

func (tds TdStruct) GetStructAnonymousStructs(tdInfo *TdInfo) []TdStruct {
	anonymousStructs := make([]TdStruct, len(tds.AnonnymousFields))
	for i, anonymousStructName := range tds.AnonnymousFields {
		anonymousStructs[i] = tdInfo.TdStructs[anonymousStructName]
	}

	// TODO: Deep copy Fields and AnonnymousFields
	return anonymousStructs
}

func ParseTdproto() (infoToFill *TdInfo, err error) {
	tdprotoFileSet := token.NewFileSet()

	infoToFill = new(TdInfo)
	infoToFill.TdEvents = make(map[string]string)
	infoToFill.TdStructs = make(map[string]TdStruct)

	tdprotoNameToAstMap, err := extractTdprotoAst(tdprotoFileSet)
	if err != nil {
		return nil, err
	}

	tdprotoAst := tdprotoNameToAstMap["tdproto"]
	for fileName, fileAst := range tdprotoAst.Files {

		basePath := path.Base(fileName)
		basePathNoExt := strings.TrimRight(basePath, path.Ext(basePath))

		err = ParseTdprotoFile(infoToFill, basePathNoExt, fileAst)
		if err != nil {
			return infoToFill, err
		}
	}

	return infoToFill, nil
}

func ParseTdprotoFile(infoToFill *TdInfo, fileName string, fileAst *ast.File) error {
	for _, declaration := range fileAst.Decls {
		switch declarationType := declaration.(type) {
		case *ast.GenDecl:
			err := ParseGenericDeclaration(infoToFill, declarationType, fileName)
			if err != nil {
				return err
			}
		case *ast.FuncDecl:
			err := parseFunctionDeclaration(infoToFill, declarationType)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func parseFunctionDeclaration(infoToFill *TdInfo, functionDeclaration *ast.FuncDecl) error {

	if !functionDeclaration.Name.IsExported() {
		return nil
	}

	if functionDeclaration.Recv == nil {
		return nil
	}

	if len(functionDeclaration.Recv.List) != 1 {
		return nil
	}

	// Only parses the GetName functions right now which maps Struct name to an event name
	if functionDeclaration.Name.Name != "GetName" {
		return nil
	}

	returnStatementAst := functionDeclaration.Body.List[0].(*ast.ReturnStmt)
	returnStatemetExpression, ok := returnStatementAst.Results[0].(*ast.BasicLit)
	if !ok {
		return nil
	}

	eventName := strings.Trim(returnStatemetExpression.Value, "\"")

	typeIdent := functionDeclaration.Recv.List[0].Type.(*ast.Ident)
	typeEventBelongsTo := typeIdent.Obj.Name

	infoToFill.TdEvents[typeEventBelongsTo] = eventName

	return nil
}

func ParseGenericDeclaration(infoToFill *TdInfo, genDeclaration *ast.GenDecl, fileName string) error {
	switch genDeclaration.Tok {
	case token.CONST:
		return parseConstDeclaration(infoToFill, genDeclaration)
	case token.TYPE:
		for _, aSpec := range genDeclaration.Specs {
			aTypeSpec := aSpec.(*ast.TypeSpec)
			err := parseTypeDeclaration(infoToFill, genDeclaration, aTypeSpec, fileName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// parse type Name struct|type {Field} declarations
func parseTypeDeclaration(infoToFill *TdInfo, genDeclaration *ast.GenDecl, declarationSpec *ast.TypeSpec, fileName string) error {

	helpString := cleanHelp(genDeclaration.Doc.Text())

	switch typeAst := declarationSpec.Type.(type) {
	case *ast.Ident:
		err := parseTypeDefinition(infoToFill, declarationSpec, typeAst)
		if err != nil {
			return err
		}
	case *ast.StructType:
		err := parseStructDefinitionInfo(infoToFill, declarationSpec, typeAst, helpString, fileName)
		if err != nil {
			return err
		}
	case *ast.ArrayType:
		err := parseArrayTypeDefinition(infoToFill, declarationSpec, typeAst)
		if err != nil {
			return err
		}
	default:
		errorLogger.Printf("WARN: Not implemented type declaration %#v", typeAst)
	}

	return nil
}

func parseArrayTypeDefinition(infoToFill *TdInfo, declarationSpec *ast.TypeSpec, arrayAst *ast.ArrayType) error {
	typeName := declarationSpec.Name.Name
	arrayExpressionAst := arrayAst.Elt.(*ast.Ident)
	arrayTypeStr := arrayExpressionAst.Name
	infoToFill.TdTypes = append(infoToFill.TdTypes, TdType{
		Name:     typeName,
		BaseType: arrayTypeStr,
		IsArray:  true,
	})
	return nil
}

func parseTypeDefinition(infoToFill *TdInfo, declarationSpec *ast.TypeSpec, typeIndent *ast.Ident) error {
	typeName := declarationSpec.Name.Name
	infoToFill.TdTypes = append(infoToFill.TdTypes, TdType{
		Name:     typeName,
		BaseType: typeIndent.Name,
	})
	return nil
}

func parseStructDefinitionInfo(infoToFill *TdInfo, declarationSpec *ast.TypeSpec, structInfo *ast.StructType, helpString string, fileName string) error {
	if helpString == "" {
		errorLogger.Printf("WARN: TdStruct missing a doc string %+v", structInfo)
		helpString = "MISSING CLASS DOCUMENTATION"
	}

	if strings.HasPrefix(strings.ToLower(helpString), "deprecated") {
		return nil
	}

	structName := declarationSpec.Name.Name
	isReadOnly := strings.Contains(helpString, "Readonly")

	var fieldsList []TdStructField
	var anonymousFieldsList []string

	for _, field := range structInfo.Fields.List {
		switch len(field.Names) {
		case 0:
			anonymousIdent := field.Type.(*ast.Ident)
			anonymousFieldName := anonymousIdent.Name
			anonymousFieldsList = append(anonymousFieldsList, anonymousFieldName)
			continue
		case 1:
		default:
			return fmt.Errorf("unexpected struct %s field name amount of %d", structName, len(field.Names))
		}

		fieldName := field.Names[0].Name
		isOmitEmpty := false
		isReadOnly := false
		isNotSerialized := false
		jsonName := fieldName
		fieldDoc := cleanHelp(field.Doc.Text())

		if field.Tag != nil {
			structTags := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))

			var jsonTags []string
			if jsonTagsStr, ok := structTags.Lookup("json"); ok {
				jsonTags = strings.Split(jsonTagsStr, ",")
			}

			for i, aTag := range jsonTags {
				if i == 0 {
					if aTag == "-" {
						isNotSerialized = true
					}

					jsonName = aTag
				} else {
					if aTag == "omitempty" {
						isOmitEmpty = true
					} else {
						return fmt.Errorf("unknown json tag %s", aTag)
					}
				}
			}

			var tdprotoTags []string
			tdprotoTagsStr, ok := structTags.Lookup("tdproto")
			if ok {
				tdprotoTags = strings.Split(tdprotoTagsStr, ",")
			}

			for _, aTag := range tdprotoTags {
				if aTag == "readonly" {
					isReadOnly = true
				} else {
					return fmt.Errorf("unknown tdproto tag %s", aTag)
				}
			}
		}

		isList := false
		isPointer := false
		fieldTypeStr := ""

		switch fieldTypeAst := field.Type.(type) {
		case *ast.Ident:
			fieldTypeStr = fieldTypeAst.Name
		case *ast.ArrayType:
			isList = true

			switch arrayTypeAst := fieldTypeAst.Elt.(type) {
			case *ast.Ident:
				fieldTypeStr = arrayTypeAst.Name
			case *ast.InterfaceType:
				fieldTypeStr = "interface{}"
			default:
				return fmt.Errorf("unknown array type %#v", arrayTypeAst)
			}

		case *ast.StarExpr:
			isPointer = true

			switch pointedType := fieldTypeAst.X.(type) {
			case *ast.Ident:
				fieldTypeStr = pointedType.Name
			case *ast.ArrayType:
				isList = true

				arrayExprAst := pointedType.Elt.(*ast.Ident)

				fieldTypeStr = arrayExprAst.Name
			case *ast.MapType:
				// TODO: Implement pointers to maps
				continue
			case *ast.SelectorExpr:
				fieldTypeStr = parseSelectorAst(pointedType)
			default:
				return fmt.Errorf("unknown pointer field of %s type %#v", structName, pointedType)
			}

		case *ast.SelectorExpr:
			fieldTypeStr = parseSelectorAst(fieldTypeAst)
		case *ast.InterfaceType:
			fieldTypeStr = "interface{}"
		default:
			return fmt.Errorf("unknown field of %s type %#v", structName, fieldTypeAst)
		}

		if fieldTypeStr == "" {
			return fmt.Errorf("empty field name %s of %s", structName, fieldName)

		}

		if fieldDoc == "" {
			fieldDoc = "DOCUMENTATION MISSING"
		}

		_, isPrimitive := GolangPrimitiveTypes[fieldTypeStr]

		fieldsList = append(fieldsList, TdStructField{
			Name:            fieldName,
			IsReadOnly:      isReadOnly,
			IsOmitEmpty:     isOmitEmpty,
			JsonName:        jsonName,
			TypeStr:         fieldTypeStr,
			IsList:          isList,
			IsPointer:       isPointer,
			IsPrimitive:     isPrimitive,
			IsNotSerialized: isNotSerialized,
			Help:            fieldDoc,
		})
	}

	sort.Slice(fieldsList, func(i, j int) bool {
		return fieldsList[i].Name < fieldsList[j].Name
	})

	infoToFill.TdStructs[structName] = TdStruct{
		Help:             helpString,
		ReadOnly:         isReadOnly,
		Name:             structName,
		Fields:           fieldsList,
		AnonnymousFields: anonymousFieldsList,
		FileName:         fileName,
	}

	return nil
}

// Parse const ( name Type = value ...) expressions
func parseConstDeclaration(infoToFill *TdInfo, genDeclaration *ast.GenDecl) error {
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

		infoToFill.TdConsts = append(infoToFill.TdConsts, TdConstFields{
			Name:  constName,
			Type:  constTypeName,
			Value: constValue.Value,
			Help:  cleanHelp(valueSpec.Doc.Text()),
		})
	}

	return nil
}

func parseSelectorAst(selectorNode *ast.SelectorExpr) string {
	expresionIdent := selectorNode.X.(*ast.Ident)
	expressionStr := expresionIdent.Name
	return expressionStr + "." + selectorNode.Sel.Name
}

func extractTdprotoAst(fileSet *token.FileSet) (map[string]*ast.Package, error) {
	tdProtoPath := tdproto.SourceDir()
	return parser.ParseDir(fileSet, tdProtoPath, nil, parser.ParseComments)
}

func cleanHelp(s string) string {
	return strings.TrimSuffix(strings.TrimSpace(strings.Join(strings.Fields(s), " ")), ".")
}

func ToSnakeCase(original string) string {
	var buildStr strings.Builder

	for i, char := range original {
		if i != 0 && unicode.IsUpper(char) {
			buildStr.WriteString("_")
		}
		buildStr.WriteString(string(unicode.ToLower(char)))
	}

	return buildStr.String()
}

func SnakeCaseToLowerCamel(original string) string {
	var buildStr strings.Builder

	nextCharToUpper := false

	for i, char := range original {
		if i != 0 && char == '_' {
			nextCharToUpper = true
			continue
		}

		nextChar := char

		if nextCharToUpper {
			nextChar = unicode.ToUpper(char)
			nextCharToUpper = false
		}

		buildStr.WriteString(string(nextChar))

	}

	return buildStr.String()
}

func LowercaseFirstLetter(original string) string {
	return strings.ToLower(original[:1]) + original[1:]
}

func UppercaseFirstLetter(original string) string {
	return strings.ToUpper(original[:1]) + original[1:]
}
