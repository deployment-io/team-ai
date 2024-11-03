package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type ContentType string

func (c ContentType) String() string { return string(c) }

func (c ContentType) EncodeBody(payload interface{}) *bytes.Buffer {
	switch c {
	case JSON:
		return encodeWithJson(payload)
	case URLENCODED:
		return encodeWithUrlEncoded(payload)
	default:
		return encodeWithJson(payload)
	}
}

func encodeWithJson(payload interface{}) *bytes.Buffer {
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(payload)
	return b
}

func encodeWithUrlEncoded(payload interface{}) *bytes.Buffer {
	pl, ok := payload.(map[string]interface{})
	if !ok {
		return bytes.NewBufferString("")
	}
	data := url.Values{}
	for k, v := range pl {
		data.Set(k, fmt.Sprintf("%v", v))
	}
	return bytes.NewBufferString(data.Encode())
}

func (h *HTTPClient) HttpPostDo(reqBodyBytes []byte, req *http.Request) (*http.Response, error) {
	for i := 0; i < h.maxRetries; i++ {
		req.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))
		var err error
		var resp *http.Response
		if h.useSlowHttp {
			resp, err = slowHTTP.Do(req)
		} else {
			resp, err = fastHTTP.Do(req)
		}
		endpoint := req.URL.String()
		if err == nil && resp.StatusCode < 500 {
			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
				if resp.Body != nil {
					respBodyBytes, writeErr := io.ReadAll(resp.Body)
					if writeErr == nil {
						log.Printf("status code for %s : %d\n", endpoint, resp.StatusCode)
						log.Println("response body: ", string(respBodyBytes))
					}
					resp.Body = io.NopCloser(bytes.NewBuffer(respBodyBytes))
					if !h.doNotLogRequestBody {
						log.Println("request body: ", string(reqBodyBytes))
					}
				}
			}
			return resp, nil
		}

		if err != nil {
			log.Println("Error:", err)
			log.Println("Retrying after 5 seconds for:", req.URL.String())
		} else {
			if resp != nil && resp.Body != nil {
				respBodyBytes, writeErr := io.ReadAll(resp.Body)
				if writeErr == nil {
					log.Println("response body: ", string(respBodyBytes))
				}
			}
			log.Println("Retrying after 5 seconds for:", req.URL.String())
			log.Println("and status code:", resp.StatusCode)
		}

		if resp != nil && resp.Body != nil {
			_, _ = io.ReadAll(resp.Body)
			closeErr := resp.Body.Close()
			if closeErr != nil {
				log.Printf("error closing response body: endpoint: %s : status: %d : error: %s\n",
					endpoint, resp.StatusCode, closeErr)
			}
		}

		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("max retry limit exceeded")
}
