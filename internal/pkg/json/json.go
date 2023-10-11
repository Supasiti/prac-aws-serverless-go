package json

import (
	"encoding/json"
)

type MarshalIndent func(v any, prefix string, indent string) ([]byte, error)

func ToJSONString(i interface{}, m ...MarshalIndent) string {
	format := json.MarshalIndent
	if len(m) > 0 {
		format = m[0]
	}
	bytes, err := format(i, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func ToInlineJSON(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}
