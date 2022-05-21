
package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/insprac/qe"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.coingecko.com/api/v3"

func get(path string, data interface{}) error {
	resp, err := http.Get(baseURL + path)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if err == nil {
			return newError("bad response (status=%d, body=%v)",
				resp.StatusCode,
				string(bodyBytes))
		} else {
			return newError("bad response (status=%d)", resp.StatusCode)
		}
	}
