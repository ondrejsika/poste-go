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

func (api PosteAPI) CreateDomain(name string) error {
	rawResp, err := api.RawCreateDomain(name)
	if err != nil {
		return err
	}
	if rawResp.StatusCode() == 201 {
		return nil
	}
	return fmt.Errorf("Err")
}

func (api PosteAPI) DeleteDomain(name string) error {
	rawResp, err := api.RawDeleteDomain(name)
	if err != nil {
		return err
	}
	if rawResp.StatusCode() == 204 {
		return nil
	}
	return fmt.Errorf("Err")
}
