package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"reflect"
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

type TdQuery struct {
	Name               string
	Help               string
	ParamsNamesAndHelp map[string]string
}

type TdStructField struct {
	Name            string
	Help            string
	JsonName        string
	SchemaName      string
	TypeStr         string
	KeyTypeStr      string
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
	Filename string
}

type TdMapType struct {
	Name         string
	Help         string
	KeyTypeStr   string
	ValueTypeStr string
	Filename     string
}

type TdPackage struct {
	TdStructs  map[string]TdStruct
	TdTypes    map[string]TdType
	TdEvents   map[string]string
	TdMapTypes map[string]TdMapType
	TdConsts   []TdConstFields
	TdQueries  map[string]TdQuery
}

type TdProto struct {
	TdForms  *TdPackage
	TdModels *TdPackage
}

type TdEnum struct {
	Name   string
	Values []string
}

func (i TdPackage) GetEnums() []TdEnum {
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

func (tds TdStruct) IsEventParams(tdInfo *TdPackage) bool {

	for eventStructName := range tdInfo.TdEvents {
		eventStruct := tdInfo.TdStructs[eventStructName]

		for _, field := range eventStruct.Fields {
			if field.Name != "Params" {
				continue
			}

			if field.TypeStr == tds.Name {
				return true
			}
		}
	}

	return false
}

func (tds TdStruct) GetStructAnonymousStructs(tdInfo *TdPackage) []TdStruct {
	anonymousStructs := make([]TdStruct, len(tds.AnonnymousFields))
	for i, anonymousStructName := range tds.AnonnymousFields {
		anonymousStructs[i] = tdInfo.TdStructs[anonymousStructName]
	}

	// TODO: Deep copy Fields and AnonnymousFields
	return anonymousStructs
}

func (tds TdStruct) GetAllJsonFields(tdInfo *TdPackage) []TdStructField {
	var allFields []TdStructField

	allFields = append(allFields, tds.Fields...)

	for _, anonStruct := range tds.GetStructAnonymousStructs(tdInfo) {
		allFields = append(allFields, anonStruct.Fields...)
	}

	return allFields
}

func ParseTdproto() (infoToFill *TdProto, err error) {
	infoToFill = new(TdProto)

	tdprotoFileSet := token.NewFileSet()

	tdModelsPackage := new(TdPackage)
	tdModelsPackage.TdEvents = make(map[string]string)
	tdModelsPackage.TdStructs = make(map[string]TdStruct)
	tdModelsPackage.TdTypes = make(map[string]TdType)
	tdModelsPackage.TdMapTypes = make(map[string]TdMapType)
	tdModelsPackage.TdQueries = make(map[string]TdQuery)

	infoToFill.TdModels = tdModelsPackage

	tdprotoNameToAstMap, err := extractTdprotoAst(tdprotoFileSet)
	if err != nil {
		return nil, err
	}

	tdprotoAst := tdprotoNameToAstMap["tdproto"]
	err = parseTdprotoAst(tdprotoAst, tdModelsPackage, nil)
	if err != nil {
		return nil, err
	}

	tdapiFileSet := token.NewFileSet()
	tdapiNameToAstMap, err := extractTdapiAst(tdapiFileSet)
	if err != nil {
		return nil, err
	}

	tdFormsPackage := new(TdPackage)
	tdFormsPackage.TdEvents = make(map[string]string)
	tdFormsPackage.TdStructs = make(map[string]TdStruct)
	tdFormsPackage.TdTypes = make(map[string]TdType)
	tdFormsPackage.TdMapTypes = make(map[string]TdMapType)

	infoToFill.TdForms = tdFormsPackage

	err = parseTdprotoAst(tdapiNameToAstMap["tdapi"], tdFormsPackage,
		&map[string]string{
			"task":         "",
			"my_reactions": "",
			"resp":         "",
			"err":          "",
		},
	)
	if err != nil {
		return nil, err
	}

	// Cherry picking
	// Task
	err = cherryPickStruct(tdModelsPackage, tdFormsPackage, "Task")
	if err != nil {
		return nil, err
	}
	// TaskFilter query
	err = cherryPickQuery(tdModelsPackage, tdFormsPackage, "TaskFilter")
	if err != nil {
		return nil, err
	}
	// MyReactions
	err = cherryPickStruct(tdModelsPackage, tdFormsPackage, "MyReactions")
	if err != nil {
		return nil, err
	}

	// Resp
	err = cherryPickStruct(tdModelsPackage, tdFormsPackage, "Resp")
	if err != nil {
		return nil, err
	}

	// Err
	err = cherryPickTypeAlias(tdModelsPackage, tdFormsPackage, "Err")
	if err != nil {
		return nil, err
	}

	return infoToFill, nil
}

func cherryPickTypeAlias(tdproto *TdPackage, tdapi *TdPackage, name string) error {

	pickObject, ok := tdapi.TdTypes[name]
	if !ok {
		return fmt.Errorf("failed to cherry pick query %s", name)
	}
	tdproto.TdTypes[name] = pickObject

	return nil
}

func cherryPickQuery(tdproto *TdPackage, tdapi *TdPackage, name string) error {

	pickObject, ok := tdapi.TdStructs[name]
	if !ok {
		return fmt.Errorf("failed to cherry pick query %s", name)
	}

	var newQuery TdQuery

	newQuery.Help = pickObject.Help
	newQuery.ParamsNamesAndHelp = make(map[string]string)
	newQuery.Name = name
	for _, field := range pickObject.Fields {
		newQuery.ParamsNamesAndHelp[field.SchemaName] = field.Help
	}

	tdproto.TdQueries[name] = newQuery

	return nil
}

func cherryPickStruct(tdproto *TdPackage, tdapi *TdPackage, name string) error {

	pickObject, ok := tdapi.TdStructs[name]
	if !ok {
		return fmt.Errorf("failed to cherry pick struct %s", name)
	}
	tdproto.TdStructs[name] = pickObject

	return nil
}

func parseTdprotoAst(packageAst *ast.Package, infoToFill *TdPackage, fileFilter *map[string]string) error {
	for fileName, fileAst := range packageAst.Files {

		basePath := path.Base(fileName)
		basePathNoExt := strings.TrimRight(basePath, path.Ext(basePath))

		if fileFilter != nil {
			_, ok := (*fileFilter)[basePathNoExt]
			if !ok {
				continue
			}
		}

		err := ParseTdprotoFile(infoToFill, basePathNoExt, fileAst)
		if err != nil {
			return err
		}
	}

	return nil
}

func ParseTdprotoFile(infoToFill *TdPackage, fileName string, fileAst *ast.File) error {
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

func parseFunctionDeclaration(infoToFill *TdPackage, functionDeclaration *ast.FuncDecl) error {

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

func ParseGenericDeclaration(infoToFill *TdPackage, genDeclaration *ast.GenDecl, fileName string) error {
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
func parseTypeDeclaration(infoToFill *TdPackage, genDeclaration *ast.GenDecl, declarationSpec *ast.TypeSpec, fileName string) error {

	helpString := cleanHelp(genDeclaration.Doc.Text())

	switch typeAst := declarationSpec.Type.(type) {
	case *ast.Ident:
		err := parseTypeDefinition(infoToFill, declarationSpec, typeAst, helpString, fileName)
		if err != nil {
			return err
		}
	case *ast.StructType:
		err := parseStructDefinitionInfo(infoToFill, declarationSpec, typeAst, helpString, fileName)
		if err != nil {
			return err
		}
	case *ast.ArrayType:
		err := parseArrayTypeDefinition(infoToFill, declarationSpec, typeAst, helpString, fileName)
		if err != nil {
			return err
		}
	case *ast.MapType:
		err := parseMapTypeDeclaration(infoToFill, declarationSpec, typeAst, helpString, fileName)
		if err != nil {
			return err
		}
	default:
		errorLogger.Printf("WARN: Not implemented type declaration %#v", typeAst)
	}

	return nil
}

func parseMapTypeDeclaration(infoToFill *TdPackage, declarationSpec *ast.TypeSpec, mapAst *ast.MapType, helpString string, fileName string) error {
	typeName := declarationSpec.Name.Name

	keyTypeStr, err := parseExprToString(mapAst.Key)
	if err != nil {
		return err
	}

	valueTypeStr, err := parseExprToString(mapAst.Value)
	if err != nil {
		return err
	}

	infoToFill.TdMapTypes[typeName] = TdMapType{
		Name:         typeName,
		KeyTypeStr:   keyTypeStr,
		ValueTypeStr: valueTypeStr,
		Help:         helpString,
		Filename:     fileName,
	}

	return nil
}

func parseArrayTypeDefinition(infoToFill *TdPackage, declarationSpec *ast.TypeSpec, arrayAst *ast.ArrayType, helpString string, fileName string) error {
	typeName := declarationSpec.Name.Name
	arrayExpressionAst := arrayAst.Elt.(*ast.Ident)
	arrayTypeStr := arrayExpressionAst.Name
	infoToFill.TdTypes[typeName] = TdType{
		Name:     typeName,
		BaseType: arrayTypeStr,
		IsArray:  true,
		Help:     helpString,
		Filename: fileName,
	}
	return nil
}

func parseTypeDefinition(infoToFill *TdPackage, declarationSpec *ast.TypeSpec, typeIndent *ast.Ident, helpString string, fileName string) error {
	typeName := declarationSpec.Name.Name
	infoToFill.TdTypes[typeName] = TdType{
		Name:     typeName,
		BaseType: typeIndent.Name,
		Help:     helpString,
		Filename: fileName,
	}
	return nil
}

func parseStructDefinitionInfo(infoToFill *TdPackage, declarationSpec *ast.TypeSpec, structInfo *ast.StructType, helpString string, fileName string) error {
	structName := declarationSpec.Name.Name

	if helpString == "" {
		errorLogger.Printf("WARN: TdStruct %s missing a doc string in file %s", structName, fileName)
	}

	if strings.HasPrefix(strings.ToLower(helpString), "deprecated") {
		return nil
	}

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
		var schemaName string

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

			var schemaTags []string
			schemaTagsStr, ok := structTags.Lookup("schema")
			if ok {
				schemaTags = strings.Split(schemaTagsStr, ",")
			}

			for _, sTag := range schemaTags {
				schemaName = sTag
			}
		}

		isList := false
		isPointer := false
		fieldTypeStr := ""
		keyTypeStr := ""

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
			case *ast.SelectorExpr:
				fieldTypeStr = parseSelectorAst(arrayTypeAst)
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
		case *ast.MapType:
			var err error
			keyTypeStr, err = parseExprToString(fieldTypeAst.Key)
			if err != nil {
				return err
			}
			fieldTypeStr, err = parseExprToString(fieldTypeAst.Value)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown field of %s type %#v", structName, fieldTypeAst)
		}

		if fieldTypeStr == "" {
			return fmt.Errorf("empty field name %s of %s", structName, fieldName)

		}

		_, isPrimitive := GolangPrimitiveTypes[fieldTypeStr]

		fieldsList = append(fieldsList, TdStructField{
			Name:            fieldName,
			IsReadOnly:      isReadOnly,
			IsOmitEmpty:     isOmitEmpty,
			JsonName:        jsonName,
			SchemaName:      schemaName,
			TypeStr:         fieldTypeStr,
			KeyTypeStr:      keyTypeStr,
			IsList:          isList,
			IsPointer:       isPointer,
			IsPrimitive:     isPrimitive,
			IsNotSerialized: isNotSerialized,
			Help:            fieldDoc,
		})
	}

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
func parseConstDeclaration(infoToFill *TdPackage, genDeclaration *ast.GenDecl) error {
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

func parseExprToString(expr interface{}) (string, error) {
	switch exprType := expr.(type) {
	case *ast.SelectorExpr:
		return parseSelectorAst(exprType), nil
	case *ast.Ident:
		return exprType.Name, nil
	case *ast.InterfaceType:
		return "interface{}", nil
	case *ast.StarExpr:
		return parseStarAst(exprType)
	}

	return "", fmt.Errorf("cannot parse expression %#v", expr)
}

func parseStarAst(starAst *ast.StarExpr) (string, error) {
	pointedType, err := parseExprToString(starAst.X)
	if err != nil {
		return "", err
	}
	return pointedType, nil
}

func parseSelectorAst(selectorNode *ast.SelectorExpr) string {
	expresionIdent := selectorNode.X.(*ast.Ident)
	expressionStr := expresionIdent.Name
	if expressionStr == "tdproto" { // HACK: when tdapi references tdproto
		return selectorNode.Sel.Name
	}
	return expressionStr + "." + selectorNode.Sel.Name
}

func extractTdprotoAst(fileSet *token.FileSet) (map[string]*ast.Package, error) {
	tdProtoPath := tdproto.SourceDir()
	return parser.ParseDir(fileSet, tdProtoPath, nil, parser.ParseComments)
}

func extractTdapiAst(fileSet *token.FileSet) (map[string]*ast.Package, error) {
	tdProtoPath := tdproto.SourceDir()
	return parser.ParseDir(fileSet, path.Join(tdProtoPath, "tdapi"), nil, parser.ParseComments)
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
