package gcg

var queriesTemplateSrc = `package {{.Package}}
// generated by github.com/emicklei/graphql-client-gen/cmd/gcg version: {{.BuildVersion }}
// DO NOT EDIT

import (
	"time"
)

var (
	_ = time.Now
)

{{- range .Queries}}

// {{.FunctionName}}Query is used for both specifying the query and capturing the response. {{.Comment}}
type {{.FunctionName}}Query struct {
	Operation string
	Errors Errors {{.ErrorsTag}}
	Data      {{if .IsArray}}[]{{end}}{{.FunctionName}}QueryData {{.DataTag}}
}

type {{.FunctionName}}QueryData struct {
	{{.ReturnType}} {{.ReturnFieldTag}}
}

// OperationName returns the actual query operation name that is used to replace "OperationName"
func (q {{.FunctionName}}Query) OperationName() string {
	return q.Operation
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q {{.FunctionName}}Query) Build(
	operationName string, // cannot be emtpy
	{{- range .Arguments}}
	{{.Name}} {{.Type}}, 
	{{- end }}
) GraphQLRequest {
	_q.Operation = operationName
	return GraphQLRequest{
		Query:         BuildQuery(_q),
		OperationName: operationName,
		Variables: map[string]interface{}{
			{{- range .Arguments}}
			"{{.JSONName}}": {{.Name}},
			{{- end }}
		},
	}
}

{{- end}}
`
