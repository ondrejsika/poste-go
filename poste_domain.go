package poste

import (
	"encoding/json"
	"fmt"
)

func (api PosteAPI) ListDomains() (*ListDomainsResponse, error) {
	rawResp, err := api.RawListDomains()
	if err != nil {
		return nil, err
	}
	respBody := rawResp.Body()
	if rawResp.StatusCode() == 200 {
		var resp ListDomainsResponse
		json.Unmarshal([]byte(respBody), &resp)
		return &resp, nil
	}
	return nil, fmt.Errorf("Err")
}
