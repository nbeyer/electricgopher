package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	original http.RoundTripper
	logger   *logrus.Logger
}

func (lrt *loggingRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	lrt.logger.Debugf("electricgopher.api: HTTP Request - %s %s", request.Method, request.URL.String())
	response, err := lrt.original.RoundTrip(request)
	if response != nil {
		lrt.logger.Debugf("electricgopher.api: HTTP Response Status - %s", response.Status)
		lrt.logger.Debugf("electricgopher.api: HTTP Response Header - %s", response.Header)
	}
	if err != nil {
		lrt.logger.Debugf("electricgopher.api: Error - %s", err.Error())
	}
	return response, err
}

func newHttpClient(logger *logrus.Logger) *http.Client {
	client := &http.Client{
		Transport: &loggingRoundTripper{
			logger:   logger,
			original: http.DefaultTransport,
		},
		Timeout: 15 * time.Second,
	}
	return client
}

func mustSerializeJson(obj interface{}) []byte {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("Unable to serialize object into JSON: %v", err))
	}
	return jsonBytes
}

func mustDeserializeJson(data []byte, obj interface{}) {
	err := json.Unmarshal(data, obj)
	if err != nil {
		panic(fmt.Sprintf("Unable to deserialize JSON into object: %v", err))
	}
}

func makeJsonPretty(data []byte) []byte {
	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, data, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("Unable to make JSON pretty: %v", err))
	}
	return prettyJson.Bytes()
}
