package kit

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"io"
)

func JsonEncoding(v interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(v)
}

func JsonDecoding(r io.Reader, v interface{}) error {
	decoder := json.NewDecoder(r)
	decoder.UseNumber()
	return decoder.Decode(v)
}

func JsonIterEncoding(v interface{}, w io.Writer) error {
	encoder := jsoniter.NewEncoder(w)
	return encoder.Encode(v)
}

func JsonIterDecoding(r io.Reader, v interface{}) error {
	decoder := jsoniter.NewDecoder(r)
	decoder.DisallowUnknownFields()
	decoder.UseNumber()
	return decoder.Decode(v)
}
