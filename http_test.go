package kit

import (
	"bytes"
	"crypto/tls"
	"testing"
	"time"
)

func TestHttpRequest(t *testing.T) {

	body := new(bytes.Buffer)
	response := new(bytes.Buffer)
	tlsCfg := &tls.Config{InsecureSkipVerify: true}

	if err := HttpRequest("POST", "http://example.com", body, response, time.Second*3, tlsCfg); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("POST", response.String())

	if err := HttpRequest("GET", "http://example.com", body, response, time.Second*3, tlsCfg); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("GET", response.String())
}

func TestHttpPost(t *testing.T) {
	body := new(bytes.Buffer)
	response := new(bytes.Buffer)
	if err := HttpPost("http://example.com", body, response); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(response.String())
}

func TestHttpGet(t *testing.T) {
	response := new(bytes.Buffer)
	if err := HttpGet("http://example.com", response); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(response.String())
}

func TestHttpsPost(t *testing.T) {
	body := new(bytes.Buffer)
	response := new(bytes.Buffer)
	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	if err := HttpsPost("http://example.com", body, response, tlsCfg); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(response.String())
}

func TestHttpsGet(t *testing.T) {
	response := new(bytes.Buffer)
	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	if err := HttpsGet("http://example.com", response, tlsCfg); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(response.String())
}

func TestHttpPostTimeout(t *testing.T) {
	body := new(bytes.Buffer)
	response := new(bytes.Buffer)
	if err := HttpPostTimeout("http://example.com", body, response, time.Second*1); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(response.String())
}

func TestHttpGetTimeout(t *testing.T) {
	response := new(bytes.Buffer)
	if err := HttpGetTimeout("http://example.com", response, time.Second*1); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(response.String())
}

func TestHttpsPostTimeout(t *testing.T) {
	body := new(bytes.Buffer)
	response := new(bytes.Buffer)
	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	if err := HttpsPostTimeout("http://example.com", body, response, time.Second*1, tlsCfg); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(response.String())
}

func TestHttpsGetTimeout(t *testing.T) {
	response := new(bytes.Buffer)
	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	if err := HttpsGetTimeout("http://example.com", response, time.Second*1, tlsCfg); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(response.String())
}
