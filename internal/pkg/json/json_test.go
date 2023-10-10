package json

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSONString(t *testing.T) {
	type args struct {
		i interface{}
	}
	type address map[string]any

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return json string",
			args: args{
				i: struct {
					Name string `json:"name"`
					Age  uint8  `json:"age"`
				}{
					Name: "Bob",
					Age:  20,
				},
			},
			want: `{"name": "Bob","age": 20}`,
		},
		{
			name: "Should omit empty fields",
			args: args{
				i: struct {
					Name string `json:"name"`
					ID   string `json:"id,omitempty"`
				}{
					Name: "Bob",
				},
			},
			want: `{"name": "Bob"}`,
		},
		{
			name: "Should properly indent nested structs",
			args: args{
				i: struct {
					Name    string  `json:"name"`
					Address address `json:"address"`
				}{
					Name: "Bob",
					Address: address{
						"state":    "VIC",
						"postcode": 3000,
					},
				},
			},
			want: `{"name": "Bob", "address": {"state": "VIC", "postcode": 3000}}`,
		},
		{
			name: "Should handle slices",
			args: args{
				i: struct {
					Name    string   `json:"name"`
					Aliases []string `json:"aliases"`
				}{
					Name: "Bob",
					Aliases: []string{
						"The builder",
						"Bob the builder",
					},
				},
			},
			want: `{"name": "Bob", "aliases": ["The builder", "Bob the builder"]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONString(tt.args.i)

			assert.JSONEqf(t, tt.want, got, "ToJSONString(%v) = %v , want %v", tt.args.i, got, tt.want)
		})
	}
}

func TestToJSONStringShouldErr(t *testing.T) {
	t.Run("Should handle when marshaller fails", func(t *testing.T) {
		input := struct {
			Name string `json:"name"`
		}{
			Name: "Bob",
		}
		mockMarshalIndent := func(v any, prefix string, indent string) ([]byte, error) {
			return []byte{}, errors.New("error message")
		}

		want := "error message"
		got := ToJSONString(input, mockMarshalIndent)

		assert.Equalf(t, want, got, "ToJSONString(%v) = %v , want %v", input, got, want)
	})
}
