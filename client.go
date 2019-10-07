package gdcli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doRequest(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%s\n", string(body))
		return nil, fmt.Errorf("doRequest()")
	}

	return res, nil
}

func generateRequestWithToken(req *http.Request, token string, secret string) {
	req.SetBasicAuth(token, secret)
	req.Header.Set("Content-Type", "application/json")
}

func decodeRequestBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}
