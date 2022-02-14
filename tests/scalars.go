package tests

// generated by github.com/emicklei/graphql-client-gen/cmd/gcg version: (devel)
// DO NOT EDIT

// These aliases are needed to model return types in Data fields
type String string
type Int int32
type Float float64
type Boolean bool

// ID can by any type, typically a string or int64
type ID interface{}

// URL is a Scalar.
type URL string

// NewURL returns a pointer to a URL value.
// Use type conversion instead e.g. v := URL(s) if you need the non-pointer value.
func NewURL(s string) *URL { v := URL(s); return &v }