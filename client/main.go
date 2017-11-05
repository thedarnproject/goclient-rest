package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/thedarnproject/goclient-rest/util"
	"github.com/thedarnproject/thedarnapi/api"
	"net/url"
	"os"
)

func main() {

	darnRESTURL := util.GetEnvVarOrDefault("DARN_REST_URL", "http://localhost")
	darnRESTPath := util.GetEnvVarOrDefault("DARN_REST_PATH", "/")
	darnPlugin := util.GetEnvVarOrDefault("DARN_PLUGIN", "bash")
	darnTrigger := util.GetEnvVarOrDefault("DARN_TRIGGER", "trap")
	darnError := util.GetEnvVarOrDefault("DARN_ERROR", "")
	darnPlatorm := util.GetEnvVarOrDefault("DARN_PLATFORM", "linux")

	if len(darnError) == 0 {
		os.Exit(1)
	}

	darnAPIValues := darn.Data{
		Plugin:   darnPlugin,
		Trigger:  darnTrigger,
		Error:    darnError,
		Platform: darnPlatorm,
	}

	u, err := url.ParseRequestURI(darnRESTURL)
	if err != nil {
		panic(err)
	}
	u.Path = darnRESTPath
	darnRESTEndpoint := u.String()

	darnRequestJSON := new(bytes.Buffer)
	json.NewEncoder(darnRequestJSON).Encode(darnAPIValues)

	hClient := &http.Client{}
	req, err := http.NewRequest("POST", darnRESTEndpoint, darnRequestJSON)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	response, err := hClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	logrus.Info(string(responseBody))
}
