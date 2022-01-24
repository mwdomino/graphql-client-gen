package gcg

var scalarsTemplateSrc = `package {{.Package}}
// generated by github.com/emicklei/graphql-client-gen/cmd/gcg on {{.Created.Format "Mon, 02 Jan 2006 15:04:05 MST" }} 
// DO NOT EDIT

{{- range .Scalars}}

// {{.Name}} is a Scalar. {{.Comment}}
type {{.Name}} string

{{- end}}
`