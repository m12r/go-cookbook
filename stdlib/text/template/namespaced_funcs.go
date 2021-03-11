package template

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

// StringsNamespace to hold all functions for usage in `Template.FuncMap` method of
// `html/template` or `text/template`.
type StringsNamespace struct{}

// singleton
var globalStringsNamespace = &StringsNamespace{}

// GetStringsNamespace returns the global strings namespace. This
// is required as template.FuncsMap expects a function pointer as
// a value.
func GetStringsNamespace() *StringsNamespace {
	return globalStringsNamespace
}

// ToUpper converts an input i to upper case string, if the input is either
// a string, slice of bytes, slice of runes or the type implements the
// fmt.Stringer interface. A nil will be converted to an empty string.
func (StringsNamespace) ToUpper(i interface{}) (interface{}, error) {
	s, err := convertToString(i)
	if err != nil {
		return nil, err
	}
	return strings.ToUpper(s), nil
}

// ToLower converts an input i to upper case string, if the input is either
// a string, slice of bytes, slice of runes or the type implements the
// fmt.Stringer interface. A nil will be converted to an empty string.
func (StringsNamespace) ToLower(i interface{}) (interface{}, error) {
	s, err := convertToString(i)
	if err != nil {
		return nil, err
	}
	return strings.ToLower(s), nil
}

// Run shows how the namespaced functions are used by the template. The
// supplied template string is parsed and then executed with the supplied
// data. The result is written to the writer. The function returns nil on
// success or an error otherwise.
func Run(w io.Writer, tmpl string, data interface{}) error {
	// New template
	tpl := template.New("test")

	// Bind function on template. Please make sure you supply
	// a function pointer as the value in the FuncMap, otherwise
	// it will panic.
	tpl.Funcs(template.FuncMap{"strings": GetStringsNamespace})

	// Parse template
	if _, err := tpl.Parse(tmpl); err != nil {
		return fmt.Errorf("cannot parse template: %w", err)
	}

	// Execute template
	if err := tpl.Execute(w, data); err != nil {
		return fmt.Errorf("cannot execute template: %w", err)
	}

	return nil
}

func convertToString(i interface{}) (string, error) {
	switch v := i.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case []rune:
		return string(v), nil
	case fmt.Stringer:
		return v.String(), nil
	case nil:
		return "", nil
	}
	return "", fmt.Errorf("cannot convert %T to string", i)
}
