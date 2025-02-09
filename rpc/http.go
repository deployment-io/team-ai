package rpc

import (
	"github.com/deployment-io/team-ai/enums/rpcs"
	"io"
	"net/http"
	"time"
)

const (
	maxIdleConnections int           = 20
	requestTimeout     time.Duration = time.Duration(10) * time.Second
	slowRequestTimeout time.Duration = time.Duration(300) * time.Second
	JSON               ContentType   = "application/json"
	URLENCODED         ContentType   = "application/x-www-form-urlencoded"
)

var fastHTTP *http.Client
var slowHTTP *http.Client

func init() {
	fastHTTP = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: requestTimeout,
	}

	slowHTTP = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: slowRequestTimeout,
	}
}

type HTTPClientInterface interface {
	Do(*http.Request) (*http.Response, error)
}

func NewHTTPClient(t rpcs.Type, useSlowHttp, doNotLogRequestBody bool, maxRetries int) *HTTPClient {
	return &HTTPClient{
		rpcType:             t,
		useSlowHttp:         useSlowHttp,
		maxRetries:          maxRetries,
		doNotLogRequestBody: doNotLogRequestBody,
	}
}

// HTTPClient represents an HTTP client for each rpc type
type HTTPClient struct {
	rpcType             rpcs.Type
	useSlowHttp         bool
	maxRetries          int
	doNotLogRequestBody bool
}

func (h *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	//can add before and after functions here
	if h.rpcType == rpcs.AzureOpenAI {
		//since azure ai langchaingo directly calls Do we'll need to handle retry etc here
		if req.Method == http.MethodPost {
			reqBodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}
			return h.HttpPostDo(reqBodyBytes, req)
		}
	}
	if h.useSlowHttp {
		return slowHTTP.Do(req)
	} else {
		return fastHTTP.Do(req)
	}
}
