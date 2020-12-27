package poste

import (
	"encoding/json"
	"fmt"
)

func (api PosteAPI) CreateBox(email string, passwordPlaintext string) error {
	rawResp, err := api.RawCreateBox(email, passwordPlaintext)
	if err != nil {
		return err
	}
	if rawResp.StatusCode() == 201 {
		return nil
	}
	return fmt.Errorf("Err")
}

func (api PosteAPI) GetBoxSieve(email string) (*BoxSieveResponse, error) {
	rawResp, err := api.RawGetBoxSieve(email)
	if err != nil {
		return nil, err
	}
	respBody := rawResp.Body()
	if rawResp.StatusCode() == 200 {
		var resp BoxSieveResponse
		json.Unmarshal([]byte(respBody), &resp)
		return &resp, nil
	}
	return nil, fmt.Errorf("Err")
}

func (api PosteAPI) UpdateBoxPassword(email string, passwordPlaintext string) error {
	rawResp, err := api.RawUpdateBoxPassword(email, passwordPlaintext)
	if err != nil {
		return err
	}
	if rawResp.StatusCode() == 204 {
		return nil
	}
	return fmt.Errorf("Err")
}

func (api PosteAPI) DeleteBox(email string) error {
	rawResp, err := api.RawDeleteBox(email)
	if err != nil {
		return err
	}
	if rawResp.StatusCode() == 204 {
		return nil
	}
	return fmt.Errorf("Err")
}
