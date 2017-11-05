package main

import (
	"fmt"
	"net/http"

	"bytes"

	"encoding/json"
	"io/ioutil"

	"github.com/thedarnproject/goclient-rest/util"
	"github.com/thedarnproject/thedarnapi/api"
)

func main() {

	darnRestEndpoint := util.GetEnvVarOrDefault("DARN_REST_ENDPOINT", "http://localhost")
	darnPlugin := util.GetEnvVarOrDefault("DARN_PLUGIN", "bash")
	darnTrigger := util.GetEnvVarOrDefault("DARN_TRIGGER", "trap")
	darnError := util.GetEnvVarOrDefault("DARN_ERROR", "")
	darnPlatorm := util.GetEnvVarOrDefault("DARN_PLATFORM", "linux")

	darnAPIValues := darn.Data{
		Plugin:   darnPlugin,
		Trigger:  darnTrigger,
		Error:    darnError,
		Platform: darnPlatorm,
	}

	darnRequestJSON := new(bytes.Buffer)
	json.NewEncoder(darnRequestJSON).Encode(darnAPIValues)

	hClient := &http.Client{}
	req, err := http.NewRequest("POST", darnRestEndpoint, darnRequestJSON)
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

	fmt.Println(string(responseBody))
}
