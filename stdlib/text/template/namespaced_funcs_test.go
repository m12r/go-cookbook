package template

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/template"
)

// Code

// Namespace to hold all functions for usage in `Template.FuncMap` method of
// `html/template` or `text/template`.
type Namespace struct{}

// ToUpper converts an input i to upper case string, if the input is either
// a string, slice of bytes, slice of runes or the type implements the
// fmt.Stringer interface. A nil will be converted to an empty string.
func (Namespace) ToUpper(i interface{}) (interface{}, error) {
	s, err := convertToString(i)
	if err != nil {
		return nil, err
	}
	return strings.ToUpper(s), nil
}

// ToLower converts an input i to upper case string, if the input is either
// a string, slice of bytes, slice of runes or the type implements the
// fmt.Stringer interface. A nil will be converted to an empty string.
func (Namespace) ToLower(i interface{}) (interface {}, error) {
	s, err := convertToString(i)
	if err != nil {
		return nil, err
	}
	return strings.ToLower(s), nil
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
	return "", fmt.Errorf("cannot convert %t to string", i)
}

// Tests

func TestNamespacedFuncs(t *testing.T) {
	testCases := []struct{
		name      string
		template  string
		data      interface{}
		expected  string
		expectErr bool
	}{
		{name:"strings.toUpper with string", template:`{{ . | strings.toUpper }}`, data:"Hello, World!", expected: "HELLO, WORLD!", expectErr: false},
		{name:"strings.toUpper with []byte", template:`{{ . | strings.toUpper }}`, data:[]byte("Hello, World!"), expected: "HELLO, WORLD!", expectErr: false},
		{name:"strings.toUpper with []rune", template:`{{ . | strings.toUpper }}`, data:[]rune("Hello, World!"), expected: "HELLO, WORLD!", expectErr: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			namespace := &Namespace{}

			tpl, err := template.New("test").
				Funcs(template.FuncMap{"strings": namespace}).
				Parse(tc.template)
			if err != nil {
				t.Fatalf("cannot create and parse template: %v", err)
			}

			buf := &bytes.Buffer{}
			execErr := tpl.Execute(buf, tc.data)
			if execErr != nil {
				if !tc.expectErr {
					t.Fatalf("unexpected error executing template: %v", execErr)
				}
				return
			}

			if tc.expected != buf.String() {
				t.Errorf("expected %v, got %v", tc.expected, buf.String())
			}
		})
	}
}