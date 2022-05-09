
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