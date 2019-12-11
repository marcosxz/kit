package kit

import (
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

const defaultHttpTimeout = time.Second * 5

func HttpRequest(method, rawUrl string, req io.Reader, rsp io.Writer, timeout time.Duration, tlsCfg *tls.Config) error {

	path, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}

	// response
	response := fasthttp.AcquireResponse()

	// request
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod(strings.ToUpper(method))
	request.Header.SetHost(path.Host)
	request.Header.SetRequestURI(path.Path)

	// read body
	if req != nil {
		if bodyBytes, err := ioutil.ReadAll(req); err != nil {
			return err
		} else {
			request.SetBody(bodyBytes)
		}
	}

	// pipeline client
	pipClient := &fasthttp.PipelineClient{
		Addr:      path.Host,
		IsTLS:     nil != tlsCfg,
		TLSConfig: tlsCfg,
	}

	if err := pipClient.DoTimeout(request, response, timeout); err != nil {
		return err
	}

	// read response
	if rsp != nil {
		if _, err := rsp.Write(response.Body()); err != nil {
			return err
		}
	}

	return nil
}

func HttpPost(rawUrl string, body io.Reader, response io.Writer) error {
	tlsCfg := &tls.Config{InsecureSkipVerify: true}
	return HttpRequest(fasthttp.MethodPost, rawUrl, body, response, defaultHttpTimeout, tlsCfg)
}

func HttpGet(rawUrl string, response io.Writer) error {
	tlsCfg := &tls.Config{InsecureSkipVerify: true}
	return HttpRequest(fasthttp.MethodGet, rawUrl, nil, response, defaultHttpTimeout, tlsCfg)
}

func HttpsPost(rawUrl string, body io.Reader, response io.Writer, tlsCfg *tls.Config) error {
	return HttpRequest(fasthttp.MethodPost, rawUrl, body, response, defaultHttpTimeout, tlsCfg)
}

func HttpsGet(rawUrl string, response io.Writer, tlsCfg *tls.Config) error {
	return HttpRequest(fasthttp.MethodGet, rawUrl, nil, response, defaultHttpTimeout, tlsCfg)
}

func HttpPostTimeout(rawUrl string, body io.Reader, response io.Writer, timeout time.Duration) error {
	tlsCfg := &tls.Config{InsecureSkipVerify: true}
	return HttpRequest(fasthttp.MethodPost, rawUrl, body, response, timeout, tlsCfg)
}

func HttpGetTimeout(rawUrl string, response io.Writer, timeout time.Duration) error {
	tlsCfg := &tls.Config{InsecureSkipVerify: true}
	return HttpRequest(fasthttp.MethodGet, rawUrl, nil, response, timeout, tlsCfg)
}

func HttpsPostTimeout(rawUrl string, body io.Reader, response io.Writer, timeout time.Duration, tlsCfg *tls.Config) error {
	return HttpRequest(fasthttp.MethodPost, rawUrl, body, response, timeout, tlsCfg)
}

func HttpsGetTimeout(rawUrl string, response io.Writer, timeout time.Duration, tlsCfg *tls.Config) error {
	return HttpRequest(fasthttp.MethodGet, rawUrl, nil, response, timeout, tlsCfg)
}
