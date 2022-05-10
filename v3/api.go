
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