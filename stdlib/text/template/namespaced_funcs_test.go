package template_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/m12r/go-cookbook/stdlib/text/template"
)

func TestNamespacedFuncsRun(t *testing.T) {
	testCases := []struct {
		name           string
		writerStringer writerStringer
		template       string
		data           interface{}
		expected       string
		expectErr      bool
	}{
		// strings.ToUpper
		{
			name:           "strings.ToUpper with string",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToUpper }}`,
			data:           "Hello, World!",
			expected:       "HELLO, WORLD!",
			expectErr:      false,
		},
		{
			name:           "strings.ToUpper with []byte",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToUpper }}`,
			data:           []byte("Hello, World!"),
			expected:       "HELLO, WORLD!",
			expectErr:      false,
		},
		{
			name:           "strings.ToUpper with []rune",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToUpper }}`,
			data:           []rune("Hello, World!"),
			expected:       "HELLO, WORLD!",
			expectErr:      false,
		},
		{
			name:           "strings.ToUpper with fmt.Stringer",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToUpper }}`,
			data:           bytes.NewBufferString("Hello, World!"),
			expected:       "HELLO, WORLD!",
			expectErr:      false,
		},
		{
			name:           "strings.ToUpper with nil",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToUpper }}`,
			data:           nil,
			expected:       "",
			expectErr:      false,
		},
		{
			name:           "strings.ToUpper cannot convert int to string",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToUpper }}`,
			data:           1,
			expected:       "",
			expectErr:      true,
		},
		{
			name:           "strings.ToUpper cannot write to writer",
			writerStringer: failingWriterStringer{},
			template:       `{{ . | strings.ToUpper }}`,
			data:           "Hello, World!",
			expected:       "",
			expectErr:      true,
		},
		// strings.ToLower
		{
			name:           "strings.ToLower with string",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToLower }}`,
			data:           "Hello, World!",
			expected:       "hello, world!",
			expectErr:      false,
		},
		{
			name:           "strings.ToLower with []byte",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToLower }}`,
			data:           []byte("Hello, World!"),
			expected:       "hello, world!",
			expectErr:      false,
		},
		{
			name:           "strings.ToLower with []rune",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToLower }}`,
			data:           []rune("Hello, World!"),
			expected:       "hello, world!",
			expectErr:      false,
		},
		{
			name:           "strings.ToLower with fmt.Stringer",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToLower }}`,
			data:           bytes.NewBufferString("Hello, World!"),
			expected:       "hello, world!",
			expectErr:      false,
		},
		{
			name:           "strings.ToLower cannot convert int to string",
			writerStringer: &bytes.Buffer{},
			template:       `{{ . | strings.ToLower }}`,
			data:           1,
			expected:       "",
			expectErr:      true,
		},
		{
			name:           "strings.ToLower cannot write to writer",
			writerStringer: failingWriterStringer{},
			template:       `{{ . | strings.ToLower }}`,
			data:           "Hello, World!",
			expected:       "",
			expectErr:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := template.Run(tc.writerStringer, tc.template, tc.data)
			if tc.expectErr {
				if err == nil {
					t.Fatal("expected error, but received none")
				}
				// success, we expected an error
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			got := tc.writerStringer.String()
			if tc.expected != got {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}

// helper types

type writerStringer interface {
	io.Writer
	fmt.Stringer
}

type failingWriterStringer struct{}

func (failingWriterStringer) Write(_ []byte) (int, error) {
	return 0, errors.New("if you see this message, your code is wrong")
}

func (failingWriterStringer) String() string {
	return ""
}
