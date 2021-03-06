package netscaler

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// GetAllStats sends a request to the Nitro API and retrieves stats for the given type.
func (c *NitroClient) GetAllStats(statsType StatsType) ([]byte, error) {
	url := c.url + "stat/" + statsType.String()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating HTTP request")
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		if resp != nil {
			io.Copy(ioutil.Discard, resp.Body)
		}
		return nil, errors.Wrap(err, "error sending request")
	}
	switch resp.StatusCode {
	case 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return body, errors.Wrap(err, "error reading response body")
		}
		return body, nil
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		return body, errors.New("read failed: " + resp.Status + " (" + string(body) + ")")
	}
}
