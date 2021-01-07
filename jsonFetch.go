package easy_fetch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type JSONError struct {
	Message string `json:"message"`
}

//FetchJSON marshals JSON structs as the request body and unmarshalls the response body into a struct.
//RequestBody and ResponseBody should be pointers to structs. Either can also be nil if not needed.
//If the server doesn't respond with JSON then ResponseBody will not attempt to be unmarshalled.
//Options can be empty if you don't need any
func FetchJSON(URL string, Method string, RequestBody interface{}, ResponseBody interface{}, Options FetchOptions) (err error) {
	// If body exists, Marshal body
	var body []byte
	if RequestBody != nil {
		body, _ = json.Marshal(RequestBody)
	}

	// Create request
	req, err := http.NewRequest(Method, URL, bytes.NewReader(body))
	if err != nil {
		return
	}

	// Add content header if necessary
	if RequestBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add options
	Options.Initiate(req)

	// Run request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	//Format response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// If we can make a response body, then make it
	if ResponseBody != nil {
		//Format JSON into an object Struct
		err = json.Unmarshal(body, &ResponseBody)
		if err != nil {
			return
		}
	}

	return
}
