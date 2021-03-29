package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/tada-team/tdproto"
)

type TadaConstFields struct {
	Name  string `json:"name"`
	Type  string
	Value string `json:"value"`
	Help  string `json:"help"`
}

type TadaStructField struct {
	Name        string
	Help        string
	JsonName    string
	TypeStr     string
	IsReadOnly  bool
	IsPointer   bool
	IsList      bool
	IsOmitEmpty bool
}

type TadaStruct struct {
	Name       string            `json:"name"`
	Help       string            `json:"help"`
	Fields     []TadaStructField `json:"fields"`
	ReadOnly   bool              `json:"readonly,omitempty"`
	EnumValues []TadaConstFields `json:"enum_values,omitempty"`
}

type TadaType struct {
	Name     string `json:"name"`
	Help     string `json:"help"`
	BaseType string
}

type TadaEvent struct {
	Name string `json:"name"`
	Help string `json:"help"`
}

func (s TadaStruct) SnakeName() string { return strcase.ToSnake(s.Name) }

func (s TadaStruct) IsEnum() bool { return len(s.EnumValues) > 0 }

type TadaInfo struct {
	TadaStructs []TadaStruct
	TadaTypes   []TadaType
	Events      []TadaEvent
	TadaConsts  []TadaConstFields
}

func ParseTdproto() (infoToFill *TadaInfo, err error) {
	tdprotoFileSet := token.NewFileSet()

	infoToFill = new(TadaInfo)

	tdprotoNameToAstMap, err := extractTdprotoAst(tdprotoFileSet)

	if err != nil {
		return nil, err
	}

	var tdprotoAst = tdprotoNameToAstMap["tdproto"]

	for fileName, fileAst := range tdprotoAst.Files {
		err = ParseTdprotoFile(infoToFill, fileName, fileAst)
		if err != nil {
			return infoToFill, err
		}
	}

	return infoToFill, nil
}

func ParseTdprotoFile(infoToFill *TadaInfo, fileName string, fileAst *ast.File) error {

	var err error = nil

	for _, declaration := range fileAst.Decls {

		switch declarationType := declaration.(type) {
		case *ast.GenDecl:
			err = ParseGenericDeclaration(infoToFill, declarationType)
		}
	}

	return err
}

func ParseGenericDeclaration(infoToFill *TadaInfo, genDeclaration *ast.GenDecl) error {
	var err error = nil
	switch genDeclaration.Tok {
	case token.CONST:
		err = parseConstDeclaration(infoToFill, genDeclaration)
	case token.TYPE:
		err = parseTypeDeclaration(infoToFill, genDeclaration)
	}

	return err
}

func parseTypeDeclaration(infoToFill *TadaInfo, genDeclaration *ast.GenDecl) error {
	// parse type Name struct|type {Field} declarations

	numberOfSpecs := len(genDeclaration.Specs)

	if numberOfSpecs != 1 {
		return fmt.Errorf("unsupported number of specs %#v", genDeclaration)
	}

	helpString := cleanHelp(genDeclaration.Doc.Text())

	declarationSpec := genDeclaration.Specs[0].(*ast.TypeSpec)

	switch typeAst := declarationSpec.Type.(type) {
	case *ast.Ident:
		err := parseTypeDefinition(infoToFill, declarationSpec, typeAst)
		if err != nil {
			return err
		}
	case *ast.StructType:
		err := parseStructDefinitioninfo(infoToFill, declarationSpec, typeAst, helpString)
		if err != nil {
			return err
		}
	default:
		fmt.Fprintf(os.Stderr, "WARN: Not implemented type declaration %#v\n", typeAst)
	}

	return nil
}

func parseTypeDefinition(infoToFill *TadaInfo, declarationSpec *ast.TypeSpec, typeIndent *ast.Ident) error {

	var typeName string = declarationSpec.Name.Name

	var newTadaType = TadaType{
		Name:     typeName,
		BaseType: typeIndent.Name,
	}

	infoToFill.TadaTypes = append(infoToFill.TadaTypes, newTadaType)
	return nil
}

func parseStructDefinitioninfo(infoToFill *TadaInfo, declarationSpec *ast.TypeSpec, structInfo *ast.StructType, helpString string) error {

	if helpString == "" {
		fmt.Fprintf(os.Stderr, "WARN: TadaStruct missing a doc string %+v\n", structInfo)
	}

	if strings.HasPrefix(strings.ToLower(helpString), "deprecated") {

		return nil
	}

	structName := declarationSpec.Name.Name

	isReadOnly := strings.Contains(helpString, "Readonly")

	var fieldsList []TadaStructField

	for _, field := range structInfo.Fields.List {
		fieldNamesLength := len(field.Names)
		if fieldNamesLength == 0 {
			// TODO: implement inheritance
			continue
		} else if fieldNamesLength == 1 {
		} else {
			return fmt.Errorf("unexpected struct %s field name amount of %d", structName, len(field.Names))
		}

		fieldName := field.Names[0].Name
		isOmitEmpty := false
		isReadOnly := false
		jsonName := fieldName

		if field.Tag != nil {
			structTags := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))

			var jsonTags []string
			jsonTagsStr, ok := structTags.Lookup("json")
			if ok {
				jsonTags = strings.Split(jsonTagsStr, ",")
			}

			for i, aTag := range jsonTags {
				if i == 0 {
					if aTag == "-" {
						continue
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
		var fieldTypeStr string

		switch fieldTypeAst := field.Type.(type) {
		case *ast.Ident:
			fieldTypeStr = fieldTypeAst.Name
		case *ast.ArrayType:
			isList = true
			arrayExprAst, ok := fieldTypeAst.Elt.(*ast.Ident)
			if !ok {
				// TODO: Implement pointers to array of interfaces
				continue
			}

			fieldTypeStr = arrayExprAst.Name
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
			case *ast.SelectorExpr:
				// TODO: Implement pointers to selectors
			default:
				panic(fmt.Sprintf("Unknown pointer field of %s type %#v", structName, pointedType))
			}

		case *ast.SelectorExpr:
			// TODO: Implement selector expresion. Example time.Time
			continue
		case *ast.InterfaceType:
			// TODO: Implement interface expression.
			continue
		default:
			return fmt.Errorf("unknown field of %s type %#v", structName, fieldTypeAst)
		}

		if fieldTypeStr == "" {
			return fmt.Errorf("empty field name %s of %s", structName, fieldName)

		}

		fieldsList = append(fieldsList, TadaStructField{
			Name:        fieldName,
			IsReadOnly:  isReadOnly,
			IsOmitEmpty: isOmitEmpty,
			JsonName:    jsonName,
			TypeStr:     fieldTypeStr,
			IsList:      isList,
			IsPointer:   isPointer,
		})
	}

	var newTadaStruct = TadaStruct{
		Help:     helpString,
		ReadOnly: isReadOnly,
		Name:     structName,
		Fields:   fieldsList,
	}

	infoToFill.TadaStructs = append(infoToFill.TadaStructs, newTadaStruct)
	return nil
}

func parseConstDeclaration(infoToFill *TadaInfo, genDeclaration *ast.GenDecl) error {
	// Parse const ( name Type = value ...) expressions

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
		if constTypeName == "" {
			fmt.Printf("WARN: const has no typeName %s\n", constName)
			continue
		}

		if len(valueSpec.Values) != 1 {
			return fmt.Errorf("expected one constant value got %+v", valueSpec.Values)
		}

		constValue, ok := valueSpec.Values[0].(*ast.BasicLit)
		if !ok {
			return fmt.Errorf("could not extract constant value %+v", valueSpec.Values[0])
		}

		infoToFill.TadaConsts = append(infoToFill.TadaConsts, TadaConstFields{
			Name:  constName,
			Type:  constTypeName,
			Value: constValue.Value,
			Help:  cleanHelp(valueSpec.Doc.Text()),
		})
	}

	return nil
}

func extractTdprotoAst(fileSet *token.FileSet) (map[string]*ast.Package, error) {
	tdProtoPath := tdproto.SourceDir()
	return parser.ParseDir(fileSet, tdProtoPath, nil, parser.ParseComments)
}

func cleanHelp(s string) string {
	return strings.TrimSuffix(strings.TrimSpace(strings.Join(strings.Fields(s), " ")), ".")
}
