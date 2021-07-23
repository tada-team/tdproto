package main

import (
	"encoding/json"
	"fmt"

	"github.com/tada-team/tdproto/codegen"
	"github.com/tada-team/tdproto/codegen/api_paths"
)

func main() {
	tdInfo, err := codegen.ParseTdproto()
	if err != nil {
		panic(err)
	}

	openapiInfo, err := generateOpenApiRoot(tdInfo.TdModels)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.MarshalIndent(openapiInfo, "", "  ")
	if err != nil {
		panic(err)
	}

	if _, err = fmt.Println(string(jsonBytes)); err != nil {
		panic(err)
	}
}

func addAllSchemas(tdInfo *codegen.TdPackage, root openApiRoot) error {
	for structName := range tdInfo.TdStructs {
		err := addStructSchema(root.Components.Schemas, structName, tdInfo)
		if err != nil {
			return err
		}
	}

	for typeName := range tdInfo.TdTypes {
		err := addTypeSchema(root.Components.Schemas, typeName, tdInfo)
		if err != nil {
			return err
		}
	}

	for mapTypeName := range tdInfo.TdMapTypes {
		err := addMapType(root.Components.Schemas, mapTypeName, tdInfo)
		if err != nil {
			return err
		}
	}

	return nil
}

func addPaths(root *openApiRoot, pathsToAdd []api_paths.PathSpec) error {

	for _, pathObject := range pathsToAdd {
		newPath := openApiPath{}

		err := pathSpecToOpenApiPath(pathObject, &newPath)
		if err != nil {
			return err
		}
		err = addPathParameters(pathObject.Path, &newPath)
		if err != nil {
			return err
		}

		root.Paths[pathObject.Path] = newPath
	}

	return nil
}

func generateOpenApiRoot(tdInfo *codegen.TdPackage) (openApiRoot, error) {
	root := openApiRoot{
		OpenApiVersion: "3.0.3",
		Info: openApiInfo{
			Title:   "tdproto",
			Version: "0.0.1", // TODO: take from git tag
		},
		Servers: []openApiServer{
			{Url: "https://web.tada.team"},
		},
		Paths: make(map[string]openApiPath),
		Security: []map[string][]string{
			{"token": {}},
		},
		Components: openApiComponents{
			Schemas: make(map[string]openApiSchema),
			SecuritySchemes: map[string]openApiSecurity{
				"token": {
					Type: openApiSecurityTypeApiKey,
					In:   openApiSecurityInHeader,
					Name: "token",
				},
			},
		},
	}

	err := addAllSchemas(tdInfo, root)
	if err != nil {
		return root, err
	}

	for _, pathCollection := range api_paths.AllPaths {
		err = addPaths(&root, pathCollection)
		if err != nil {
			return root, err
		}
	}

	return root, nil
}
