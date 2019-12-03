package kit

import (
	"bytes"
	"testing"
)

type testJson struct {
	A int                    `json:"a"`
	B bool                   `json:"b"`
	C string                 `json:"c"`
	D float64                `json:"d"`
	E map[string]interface{} `json:"e"`
	F *testJson              `json:"f"`
	G int
}

var testJsonInstance = &testJson{
	A: 4545768765467687,
	B: false,
	C: "testJson",
	D: 1.4545768765467687,
	E: map[string]interface{}{
		"map1": 1,
		"map2": true,
		"map3": "textJson",
		"map4": 1.4545768765467687,
	},
	F: &testJson{
		A: 4545768765467687,
		B: false,
		C: "testJson",
		D: 1.4545768765467687,
		E: map[string]interface{}{
			"map1": 1,
			"map2": true,
			"map3": "textJson",
			"map4": 1.4545768765467687,
		},
	},
}

func TestJson(t *testing.T) {
	var buffer = new(bytes.Buffer)
	if err := JsonEncoding(testJsonInstance, buffer); err != nil {
		t.Error(err)
		t.FailNow()
	}
	var decodingTest = &testJson{}
	if err := JsonDecoding(buffer, &decodingTest); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", decodingTest)
}

func TestJsonIter(t *testing.T) {
	var buffer = new(bytes.Buffer)
	if err := JsonIterEncoding(testJsonInstance, buffer); err != nil {
		t.Error(err)
		t.FailNow()
	}
	var decodingTest = &testJson{}
	if err := JsonIterDecoding(buffer, &decodingTest); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", decodingTest)
}
