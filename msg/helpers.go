package msg

import "encoding/json"

func DecodeInt(args json.RawMessage) (int, error) {
	var i int

	return i, json.Unmarshal(args, &i)
}

func DecodeString(args json.RawMessage) (string, error) {
	var s string

	return s, json.Unmarshal(args, &s)
}

func DecodeStruct(args json.RawMessage, i interface{}) error {
	return json.Unmarshal(args, i)
}
