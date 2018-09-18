package configmanager

import (
	"encoding/json"
	"io"
)

//supported encoding names
const (
	JSON = "json"
)

// JSONDecode decodes json
func JSONDecode(reader io.Reader, data interface{}) error {
	return json.NewDecoder(reader).Decode(data)
}

// GetDecoder returns decoder given encoding name
// if the decoder string is not idenfied use JSON
func GetDecoder(decoder string) func(io.Reader, interface{}) error {
	switch decoder {
	case JSON:
		return JSONDecode
	default:
		return JSONDecode
	}
}
