package codegen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/tada-team/tdproto"
)

type EnumValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Help  string `json:"help"`
}

func (e EnumValue) DartValue() string { return strings.ReplaceAll(e.Value, `"`, `'`) }

func (e EnumValue) DartName() string { return strcase.ToLowerCamel(e.Name) }

type TadaField struct {
	Name         string `json:"name"`
	JSName       string `json:"js_name"`
	Help         string `json:"help"`
	Json         string `json:"json"`
	Type         string `json:"type"`
	Readonly     bool   `json:"readonly"`
	Null         bool   `json:"null,omitempty"`
	List         bool   `json:"list,omitempty"`
	InternalType bool   `json:"internal_type,omitempty"`
	Omitempty    bool   `json:"omitempty,omitempty"`
	TSType       string `json:"ts_type"`
	DartType     string `json:"dart_type"`
	TSDefault    string `json:"ts_default"`
}

func (f TadaField) DartRequired() bool {
	return !(f.Null || f.Omitempty)
}

func (f TadaField) DartName() string {
	switch f.JSName {
	case "default":
		return "isDefault"
	default:
		return f.JSName
	}
}

type TadaStruct struct {
	Name       string       `json:"name"`
	Help       string       `json:"help"`
	Fields     []*TadaField `json:"fields"`
	Readonly   bool         `json:"readonly,omitempty"`
	EnumValues []EnumValue  `json:"enum_values,omitempty"`
}

type TadaEvent struct {
	Name string `json:"name"`
	Help string `json:"help"`
}

func (s TadaStruct) SnakeName() string { return strcase.ToSnake(s.Name) }

func (s TadaStruct) IsEnum() bool { return len(s.EnumValues) > 0 }

type TadaInfo struct {
	TadaStructs []TadaStruct
	Events      []TadaEvent
}

func Render(wr io.Writer, s string) error {
	tpl, err := template.New("").Parse(s)
	if err != nil {
		return err
	}

	p, err := ParseTdProto()
	if err != nil {
		return err
	}

	return tpl.Execute(wr, p)
}

func ParseTdProto() (p TadaInfo, err error) {
	fset := token.NewFileSet()
	enumsMap := make(map[string][]EnumValue)
	if err := extractTdprotoAst(fset, token.CONST, func(gen *ast.GenDecl, eventNames map[string]*ast.FuncDecl) error {
		for _, spec := range gen.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || valueSpec.Type == nil {
				continue
			}

			typeName := fmt.Sprintf("%s", valueSpec.Type) // XXX:
			if typeName == "" {
				continue
			}

			name := valueSpec.Names[0]
			val, ok := valueSpec.Values[0].(*ast.BasicLit)
			if !ok {
				continue
			}

			if enumsMap[typeName] == nil {
				enumsMap[typeName] = []EnumValue{}
			}

			enumsMap[typeName] = append(enumsMap[typeName], EnumValue{
				Name:  name.Name,
				Value: val.Value,
				Help:  cleanHelp(valueSpec.Doc.Text()),
			})
		}
		return nil
	}); err != nil {
		return p, err
	}

	types := make(map[string]struct{})
	if err := extractTdprotoAst(fset, token.TYPE, func(gen *ast.GenDecl, eventNames map[string]*ast.FuncDecl) error {
		s := TadaStruct{Help: cleanHelp(gen.Doc.Text())}
		if s.Help == "" || strings.HasPrefix(strings.ToLower(s.Help), "deprecated") {
			return nil
		}

		s.Readonly = strings.Contains(s.Help, "Readonly") // XXX
		for _, spec := range gen.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			s.Name = typeSpec.Name.Name
			if eventNames[s.Name] != nil {
				p.Events = append(p.Events, TadaEvent{
					Name: strcase.ToDelimited(s.Name, '.'),
					Help: s.Help,
				})
				continue
			}

			types[s.Name] = struct{}{}
			s.EnumValues = enumsMap[s.Name]
			if len(s.EnumValues) > 0 {
				for i := range s.EnumValues {
					s.EnumValues[i].Name = strings.TrimSuffix(s.EnumValues[i].Name, s.Name)
				}
			} else {
				switch typeSpec.Type.(type) {
				case *ast.StructType:

					st := typeSpec.Type.(*ast.StructType)
					for _, field := range st.Fields.List {
						name := field.Names[0].Name

						omitempty := false
						readonly := false
						jsonName := name
						if field.Tag != nil {
							tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
							jsontag := strings.Split(tag.Get("json"), ",")

							omitempty = len(jsontag) == 2 && jsontag[1] == "omitempty"

							jsonName = jsontag[0]
							if jsonName == "-" {
								continue
							}

							for _, tag := range strings.Split(tag.Get("tdproto"), ",") {
								if tag == "readonly" {
									readonly = true
									break
								}
							}
						}

						var typeNameBuf bytes.Buffer
						if err := printer.Fprint(&typeNameBuf, fset, field.Type); err != nil {
							return err
						}

						nullable := false
						list := false

						typeDescr := typeNameBuf.String()
						if strings.HasPrefix(typeDescr, "*") {
							nullable = true
							typeDescr = typeDescr[1:]
						}

						if strings.HasPrefix(typeDescr, "[]") {
							list = true
							typeDescr = typeDescr[2:]
						}

						tsType := typesMap[typeDescr].TypeScript
						if tsType == "" {
							tsType = typeDescr
						}

						dartType := typesMap[typeDescr].Dart
						if dartType == "" {
							dartType = typeDescr
						}

						if nullable && !omitempty && typeDescr == "JID" { // XXX:
							nullable = false
						}

						help := cleanHelp(field.Doc.Text())
						if help == "" {
							help = name
						}

						tsDefault := ""
						if omitempty {
							if nullable {
								tsDefault = "null"
							} else {
								tsDefault = tsDefaultMap[tsType]
								if tsDefault == "" {
									tsDefault = "unknown: " + tsType
								}
							}
						}

						s.Fields = append(s.Fields, &TadaField{
							Name:      name,
							Help:      help,
							JSName:    strings.ToLower(name[:1]) + name[1:],
							TSType:    tsType,
							TSDefault: tsDefault,
							DartType:  dartType,
							Json:      jsonName,
							Type:      typeDescr,
							Null:      nullable,
							List:      list,
							Omitempty: omitempty,
							Readonly:  readonly,
						})
					}
				}
			}

			p.TadaStructs = append(p.TadaStructs, s)
		}
		return nil
	}); err != nil {
		return p, err
	}

	sort.Slice(p.TadaStructs, func(i, j int) bool {
		return p.TadaStructs[i].Name < p.TadaStructs[j].Name
	})

	sort.Slice(p.Events, func(i, j int) bool {
		return p.Events[i].Name < p.Events[j].Name
	})

	for _, s := range p.TadaStructs {
		for _, f := range s.Fields {
			_, f.InternalType = types[f.Type]
		}
	}

	return p, nil
}

func extractTdprotoAst(fset *token.FileSet, tok token.Token, fn func(gen *ast.GenDecl, eventNames map[string]*ast.FuncDecl) error) error {
	tdProtoPath := tdproto.SourceDir()
	_, err := parser.ParseDir(fset, tdProtoPath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	return nil
}

func cleanHelp(s string) string {
	return strings.TrimSuffix(strings.TrimSpace(strings.Join(strings.Fields(s), " ")), ".")
}
