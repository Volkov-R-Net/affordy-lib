package httpreq

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// PostReq - Use for POST requests
func PostReq(data map[string]interface{}, url string, headers ...map[string]string) (response map[string]interface{}, err error) {

	reqBody, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	if len(headers) != 0 {
		for i := 0; i < len(headers); i++ {
			for s, e := range headers[i] {
				req.Header.Add(s, e)
			}
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
