package main

type openApiRef struct {
	Ref string `json:"$ref,omitempty"`
}

type openApiServer struct {
	Url         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type openApiInfo struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version"`
}

type openApiIn string

const (
	InQuery  openApiIn = "query"
	InHeader openApiIn = "header"
	InPath   openApiIn = "path"
	InCookie openApiIn = "cookie"
)

type openApiParameter struct {
	Name        string        `json:"name"`
	In          openApiIn     `json:"in,omitempty"`
	Required    bool          `json:"required,omitempty"`
	Description string        `json:"description,omitempty"`
	Schema      openApiSchema `json:"schema"`
}

type openApiContent struct {
	Schema openApiSchema `json:"schema"`
}

type openApiContents struct {
	ApplicationJSON               *openApiContent `json:"application/json,omitempty"`
	ApplicationXWWWFormUrlencoded *openApiContent `json:"application/x-www-form-urlencoded,omitempty"`
}

type openApiRequestBody struct {
	Content openApiContents `json:"content,omitempty"`
}

type openApiResponse struct {
	Description string          `json:"description"`
	Content     openApiContents `json:"content"`
}

type openApiOperation struct {
	Tags        []string                   `json:"tags,omitempty"`
	Summary     string                     `json:"summary,omitempty"`
	Description string                     `json:"description,omitempty"`
	Responses   map[string]openApiResponse `json:"responses,omitempty"`
	RequestBody *openApiRequestBody        `json:"requestBody,omitempty"`
	Security    []map[string][]string      `json:"security,omitempty"`
	Parameters  []openApiParameter         `json:"parameters,omitempty"`
}

type openApiPath struct {
	Parameters []openApiParameter `json:"parameters,omitempty"`
	Summary    string             `json:"summary,omitempty"`
	Get        *openApiOperation  `json:"get,omitempty"`
	Post       *openApiOperation  `json:"post,omitempty"`
	Put        *openApiOperation  `json:"put,omitempty"`
	Delete     *openApiOperation  `json:"delete,omitempty"`
}

type openApiType string

const (
	openApiString  openApiType = "string"
	openApiArray   openApiType = "array"
	openApiObject  openApiType = "object"
	openApiBoolean openApiType = "boolean"
	openApiNumber  openApiType = "number"
	openApiInteger openApiType = "integer"
)

type openApiFormat string

const (
	openApiUuid     openApiFormat = "uuid"
	openApiBinary   openApiFormat = "binary"
	openApiEmail    openApiFormat = "email"
	openApiInt32    openApiFormat = "int32"
	openApiInt64    openApiFormat = "int64"
	openApiBase64   openApiFormat = "base64"
	openApiDateTime openApiFormat = "date-time"
)

type openApiSchema struct {
	openApiRef
	Type                 openApiType              `json:"type,omitempty"`
	Format               openApiFormat            `json:"format,omitempty"`
	IsNullable           bool                     `json:"nullable,omitempty"`
	Properties           map[string]openApiSchema `json:"properties,omitempty"`
	Items                *openApiSchema           `json:"items,omitempty"`
	Required             []string                 `json:"required,omitempty"`
	Description          string                   `json:"description,omitempty"`
	Example              interface{}              `json:"example,omitempty"`
	AdditionalProperties *openApiSchema           `json:"additionalProperties,omitempty"`
}

type openApiComponents struct {
	Schemas         map[string]openApiSchema   `json:"schemas,omitempty"`
	SecuritySchemes map[string]openApiSecurity `json:"securitySchemes,omitempty"`
}

type openApiSecurityType string

const (
	openApiSecurityTypeApiKey        openApiSecurityType = "apiKey"
	openApiSecurityTypeHttp          openApiSecurityType = "http"
	openApiSecurityTypeO2Auth        openApiSecurityType = "oauth2"
	openApiSecurityTypeOpenIdConnect openApiSecurityType = "openIdConnect"
)

type openApiSecurityIn string

const (
	openApiSecurityInQuery  openApiSecurityIn = "query"
	openApiSecurityInHeader openApiSecurityIn = "header"
	openApiSecurityInCookie openApiSecurityIn = "cookie"
)

type openApiSecurity struct {
	Type openApiSecurityType `json:"type"`
	Name string              `json:"name"`
	In   openApiSecurityIn   `json:"in"`
	// TODO: add the rest of the field if they are ever requered
}

type openApiRoot struct {
	OpenApiVersion string                 `json:"openapi"`
	Info           openApiInfo            `json:"info"`
	Servers        []openApiServer        `json:"servers,omitempty"`
	Paths          map[string]openApiPath `json:"paths"`
	Components     openApiComponents      `json:"components"`
	Security       []map[string][]string  `json:"security,omitempty"`
	Tags           []string               `json:"tags,omitempty"`
}
